package com.lomolo.copodapp.http

import okhttp3.OkHttpClient
import java.util.concurrent.TimeUnit

fun getHttpClient(): OkHttpClient {
    return OkHttpClient.Builder()
        .connectTimeout(5, TimeUnit.MINUTES)
        .callTimeout(5, TimeUnit.MINUTES)
        .readTimeout(5, TimeUnit.MINUTES)
        .writeTimeout(5, TimeUnit.MINUTES)
        .build()
}