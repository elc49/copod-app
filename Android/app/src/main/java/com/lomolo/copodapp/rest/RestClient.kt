package com.lomolo.copodapp.rest

import com.lomolo.copodapp.network.IRestFul
import retrofit2.Retrofit

fun getRestService(retrofit: Retrofit): IRestFul {
    return retrofit.create(IRestFul::class.java)
}