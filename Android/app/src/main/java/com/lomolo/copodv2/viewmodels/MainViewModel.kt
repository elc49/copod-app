package com.lomolo.copodv2.viewmodels

import android.net.Uri
import android.util.Log
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
                Log.d(_tag, isLoggedIn.toString())
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
                Log.d(_tag, e.message ?: "Something went wrong")
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
        private val _tag = "MainViewModel"
    }

    fun initialize() {
        viewModelScope.launch {
            web3Auth.initialize().await()
            isUserLoggedIn()
        }
    }
}