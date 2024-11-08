package com.lomolo.copodapp.graphql

import com.apollographql.apollo.ApolloClient
import com.apollographql.apollo.cache.normalized.normalizedCache
import com.apollographql.apollo.cache.normalized.sql.SqlNormalizedCacheFactory
import com.apollographql.apollo.network.okHttpClient
import com.lomolo.copodapp.BuildConfig
import okhttp3.OkHttpClient

fun getGraphqlClient(httpClient: OkHttpClient): ApolloClient {
    val sqlCacheStore = SqlNormalizedCacheFactory("graphql.db")
    val baseApi = BuildConfig.apilocal
    return ApolloClient.Builder()
        .okHttpClient(okHttpClient = httpClient)
        .httpServerUrl("${baseApi}/graphql")
        .normalizedCache(sqlCacheStore)
        .build()
}