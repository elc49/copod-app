package com.lomolo.copodapp.ui.viewmodels

import android.net.Uri
import android.util.Log
import androidx.compose.runtime.getValue
import androidx.compose.runtime.mutableStateOf
import androidx.compose.runtime.setValue
import androidx.lifecycle.ViewModel
import androidx.lifecycle.viewModelScope
import com.apollographql.apollo.ApolloClient
import com.apollographql.apollo.cache.normalized.apolloStore
import com.lomolo.copodapp.model.DeviceDetails
import com.lomolo.copodapp.network.IRestFul
import com.lomolo.copodapp.repository.IWeb3Auth
import com.web3auth.core.types.LoginParams
import com.web3auth.core.types.Provider
import com.web3auth.core.types.UserInfo
import kotlinx.coroutines.flow.MutableStateFlow
import kotlinx.coroutines.flow.StateFlow
import kotlinx.coroutines.flow.asStateFlow
import kotlinx.coroutines.flow.update
import kotlinx.coroutines.future.await
import kotlinx.coroutines.launch
import org.web3j.crypto.Credentials

interface InitializeSdk {
    data object Success : InitializeSdk
    data object Loading : InitializeSdk
    data class Error(val msg: String?) : InitializeSdk
}

interface LoginSdk {
    data object Success : LoginSdk
    data object Loading : LoginSdk
    data class Error(val msg: String?) : LoginSdk
}

interface GetDeviceDetails {
    data object Success : GetDeviceDetails
    data object Loading : GetDeviceDetails
    data class Error(val msg: String?) : GetDeviceDetails
}

class MainViewModel(
    private val web3Auth: IWeb3Auth,
    private val restApiService: IRestFul,
    private val apolloClient: ApolloClient,
) : ViewModel() {
    var credentials: Credentials? by mutableStateOf(null)
        private set

    var userInfo: UserInfo? by mutableStateOf(null)
        private set

    private val _isLoggedIn: MutableStateFlow<Boolean> = MutableStateFlow(false)
    val isLoggedIn: StateFlow<Boolean> = _isLoggedIn.asStateFlow()

    var initializeSdk: InitializeSdk by mutableStateOf(InitializeSdk.Success)
        private set

    var loginSdk: LoginSdk by mutableStateOf(LoginSdk.Success)
        private set

    private val _deviceDetails: MutableStateFlow<DeviceDetails> = MutableStateFlow(DeviceDetails())
    val deviceDetails: StateFlow<DeviceDetails> = _deviceDetails.asStateFlow()

    var gettingDeviceDetails: GetDeviceDetails by mutableStateOf(GetDeviceDetails.Success)
        private set

    private fun prepareCredentials() {
        credentials = web3Auth.getCredentials(privateKey())
    }

    private fun prepareUserInfo() {
        userInfo = web3Auth.getUserInfo()
    }

    private fun privateKey(): String {
        return web3Auth.getPrivateKey()
    }

    private fun isUserLoggedIn() {
        viewModelScope.launch {
            try {
                val isLoggedIn = web3Auth.isAuthenticated()
                if (isLoggedIn) {
                    prepareCredentials()
                    prepareUserInfo()
                }
                _isLoggedIn.emit(isLoggedIn)
            } catch (e: Exception) {
                _isLoggedIn.emit(false)
                e.printStackTrace()
            }
        }
    }

    fun login() {
        if (loginSdk !is LoginSdk.Loading) {
            loginSdk = LoginSdk.Loading
            val loginParams = LoginParams(
                loginProvider = Provider.GOOGLE,
            )
            viewModelScope.launch {
                loginSdk = try {
                    web3Auth.login(loginParams).await()
                    prepareCredentials()
                    prepareUserInfo()
                    _isLoggedIn.emit(true)
                    LoginSdk.Success
                } catch (e: Exception) {
                    Log.d(TAG, e.message ?: "Something went wrong")
                    _isLoggedIn.emit(false)
                    LoginSdk.Error(e.message)
                }
            }
        }
    }

    fun logOut() {
        viewModelScope.launch {
            try {
                web3Auth.logout().await()
                initialize()
                apolloClient.apolloStore.clearAll()
            } catch (e: Exception) {
                e.printStackTrace()
            } finally {
                _isLoggedIn.emit(false)
            }
        }
    }

    fun setResultUrl(uri: Uri?) {
        viewModelScope.launch {
            web3Auth.setResultUri(uri)
        }
    }

    companion object {
        private const val TAG = "MainViewModel"
    }

    fun initialize() {
        initializeSdk = InitializeSdk.Loading
        viewModelScope.launch {
            initializeSdk = try {
                web3Auth.initialize().await()
                isUserLoggedIn()
                InitializeSdk.Success
            } catch (e: Exception) {
                e.printStackTrace()
                Log.d(TAG, e.message ?: "Something went wrong")
                logOut()
                InitializeSdk.Error(e.message)
            }
        }
    }

    fun getDeviceDetails() {
        if (gettingDeviceDetails !is GetDeviceDetails.Loading) {
            gettingDeviceDetails = GetDeviceDetails.Loading
            viewModelScope.launch {
                gettingDeviceDetails = try {
                    val res = restApiService.getIpinfo()
                    _deviceDetails.update {
                        res
                    }
                    GetDeviceDetails.Success
                } catch (e: Exception) {
                    Log.d(TAG, e.message ?: "Something went wrong")
                    GetDeviceDetails.Error(e.message)
                }
            }
        }
    }

    init {
        getDeviceDetails()
    }
}