package com.lomolo.copodapp.rest

import com.lomolo.copodapp.network.RestFul
import retrofit2.Retrofit

fun getRestService(retrofit: Retrofit): RestFul {
    return retrofit.create(RestFul::class.java)
}