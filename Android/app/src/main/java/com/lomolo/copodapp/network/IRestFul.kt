package com.lomolo.copodapp.network

import com.lomolo.copodapp.model.DeviceDetails
import okhttp3.MultipartBody
import retrofit2.http.GET
import retrofit2.http.Headers
import retrofit2.http.Multipart
import retrofit2.http.POST
import retrofit2.http.Part

data class UploadRes(
    val imageUri: String,
)

interface IRestFul {
    @Multipart
    @POST("/api/upload")
    suspend fun uploadDoc(@Part body: MultipartBody.Part): UploadRes
    @Headers("Content-Type: application/json")
    @GET("/api/ipinfo")
    suspend fun getIpinfo(): DeviceDetails
}