package com.lomolo.copodv2.viewmodels

import android.net.Uri
import android.util.Log
import androidx.compose.runtime.getValue
import androidx.compose.runtime.mutableStateOf
import androidx.compose.runtime.setValue
import androidx.lifecycle.ViewModel
import androidx.lifecycle.viewModelScope
import com.lomolo.copodv2.repository.IWeb3Auth
import com.web3auth.core.types.ExtraLoginOptions
import com.web3auth.core.types.LoginParams
import com.web3auth.core.types.Provider
import com.web3auth.core.types.UserInfo
import kotlinx.coroutines.flow.MutableStateFlow
import kotlinx.coroutines.flow.StateFlow
import kotlinx.coroutines.flow.asStateFlow
import kotlinx.coroutines.future.await
import kotlinx.coroutines.launch
import org.web3j.crypto.Credentials

class MainViewModel(
    private val web3Auth: IWeb3Auth
) : ViewModel() {
    lateinit var credentials: Credentials
    lateinit var userInfo: UserInfo
    private val _isLoggedIn: MutableStateFlow<Boolean> = MutableStateFlow(false)
    val isLoggedIn: StateFlow<Boolean> = _isLoggedIn.asStateFlow()

    var initializeSdk: InitializeSdk by mutableStateOf(InitializeSdk.Success)
        private set

    private fun prepareCredentials() {
        credentials = Credentials.create(privateKey())
        println(credentials)
    }

    private fun prepareUserInfo() {
        userInfo = web3Auth.getUserInfo()
        println(userInfo)
    }

    private fun privateKey(): String {
        return web3Auth.getPrivateKey()
    }

    private fun isUserLoggedIn() {
        viewModelScope.launch {
            try {
                val isLoggedIn = web3Auth.isAuthenticated()
                Log.d(TAG, isLoggedIn.toString())
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
        val loginParams = LoginParams(
            loginProvider = Provider.EMAIL_PASSWORDLESS,
            extraLoginOptions = ExtraLoginOptions(login_hint = "workockmoses@gmail.com")
        )
        viewModelScope.launch {
            try {
                web3Auth.login(loginParams).await()
                prepareCredentials()
                prepareUserInfo()
                _isLoggedIn.emit(true)
            } catch (e: Exception) {
                Log.d(TAG, e.message ?: "Something went wrong")
                _isLoggedIn.emit(false)
            }
        }
    }

    fun logOut() {
        viewModelScope.launch {
            try {
                web3Auth.logout().await()
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
            try {
                web3Auth.initialize().await()
                isUserLoggedIn()
                initializeSdk = InitializeSdk.Success
            } catch (e: Exception) {
                e.printStackTrace()
                Log.d(TAG, e.message ?: "Something went wrong")
                initializeSdk = InitializeSdk.Error(e.message)
            }
        }
    }
}

interface InitializeSdk {
    data object Success: InitializeSdk
    data object Loading: InitializeSdk
    data class Error(val msg: String?): InitializeSdk
}