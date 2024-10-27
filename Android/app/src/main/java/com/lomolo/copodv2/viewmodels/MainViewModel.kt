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
import kotlinx.coroutines.flow.update
import kotlinx.coroutines.future.await
import kotlinx.coroutines.launch
import org.web3j.crypto.Credentials

interface InitializeSdk {
    data object Success: InitializeSdk
    data object Loading: InitializeSdk
    data class Error(val msg: String?): InitializeSdk
}

interface LoginSdk {
    data object Success: LoginSdk
    data object Loading: LoginSdk
    data class Error(val msg: String?): LoginSdk
}

data class LoginInput(
    val email: String = "",
)

class MainViewModel(
    private val web3Auth: IWeb3Auth
) : ViewModel() {
    lateinit var credentials: Credentials
    lateinit var userInfo: UserInfo
        private set
    private val _isLoggedIn: MutableStateFlow<Boolean> = MutableStateFlow(false)
    val isLoggedIn: StateFlow<Boolean> = _isLoggedIn.asStateFlow()

    var initializeSdk: InitializeSdk by mutableStateOf(InitializeSdk.Success)
        private set
    var loginSdk: LoginSdk by mutableStateOf(LoginSdk.Success)
        private set

    private val _loginInput: MutableStateFlow<LoginInput> = MutableStateFlow(LoginInput())
    val loginInput: StateFlow<LoginInput> = _loginInput.asStateFlow()

    fun setEmail(email: String) {
        _loginInput.update {
            it.copy(email = email)
        }
    }

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

    fun isValidEmail(email: String): Boolean {
        val emailRegex = "^[A-Za-z0-9+_.-]+@[A-Za-z0-9.-]+\\.[A-Za-z]{2,6}$"
        return email.matches(emailRegex.toRegex())
    }

    fun login() {
        val email = _loginInput.value.email
        if (loginSdk !is LoginSdk.Loading && isValidEmail(email)) {
            loginSdk = LoginSdk.Loading
            val loginParams = LoginParams(
                loginProvider = Provider.EMAIL_PASSWORDLESS,
                extraLoginOptions = ExtraLoginOptions(login_hint = email)
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