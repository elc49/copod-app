package com.lomolo.copodapp.json

import com.squareup.moshi.Moshi
import com.squareup.moshi.kotlin.reflect.KotlinJsonAdapterFactory

fun getJsonAdapter(): Moshi {
    return Moshi.Builder()
        .add(KotlinJsonAdapterFactory())
        .build()
}