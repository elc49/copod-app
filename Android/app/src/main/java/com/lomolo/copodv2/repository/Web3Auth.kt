package com.lomolo.copodv2.repository

import android.net.Uri
import com.web3auth.core.Web3Auth
import com.web3auth.core.types.LoginParams
import com.web3auth.core.types.UserInfo
import com.web3auth.core.types.Web3AuthResponse
import java.util.concurrent.CompletableFuture

interface IWeb3Auth {
    suspend fun login(loginParams: LoginParams): CompletableFuture<Web3AuthResponse>
    suspend fun logout(): CompletableFuture<Void>
    fun getUserInfo(): UserInfo
    fun getPrivateKey(): String
    suspend fun initialize(): CompletableFuture<Void>
    suspend fun setResultUri(uri: Uri?)
    suspend fun isAuthenticated(): Boolean
}

class Web3Auth(
    private val web3Auth: Web3Auth
) : IWeb3Auth {
    override suspend fun login(loginParams: LoginParams) = web3Auth.login(loginParams)
    override suspend fun logout() = web3Auth.logout()
    override fun getPrivateKey() = web3Auth.getPrivkey()
    override fun getUserInfo(): UserInfo {
        try {
            return web3Auth.getUserInfo()!!
        } catch (e: Exception) {
            throw e
        }
    }

    override suspend fun initialize() = web3Auth.initialize()
    override suspend fun setResultUri(uri: Uri?) = web3Auth.setResultUrl(uri)
    override suspend fun isAuthenticated() = web3Auth.getPrivkey().isNotEmpty()
}