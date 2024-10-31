package com.lomolo.copodapp.retrofit

import com.lomolo.copodapp.BuildConfig
import com.squareup.moshi.Moshi
import okhttp3.OkHttpClient
import retrofit2.Retrofit
import retrofit2.converter.moshi.MoshiConverterFactory

fun getRestApiClient(okHttpClient: OkHttpClient, moshi: Moshi): Retrofit {
    return Retrofit.Builder()
        .baseUrl(BuildConfig.apilocal)
        .addConverterFactory(MoshiConverterFactory.create(moshi))
        .client(okHttpClient)
        .build()
}