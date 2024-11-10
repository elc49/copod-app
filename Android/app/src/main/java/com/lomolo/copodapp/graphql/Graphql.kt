package com.lomolo.copodapp.graphql

import com.apollographql.apollo.ApolloClient
import com.apollographql.apollo.cache.normalized.normalizedCache
import com.apollographql.apollo.cache.normalized.sql.SqlNormalizedCacheFactory
import com.apollographql.apollo.network.okHttpClient
import com.apollographql.apollo.network.ws.GraphQLWsProtocol
import com.lomolo.copodapp.BuildConfig
import kotlinx.coroutines.CoroutineScope
import kotlinx.coroutines.Dispatchers
import kotlinx.coroutines.delay
import kotlinx.coroutines.launch
import okhttp3.OkHttpClient

fun getGraphqlClient(httpClient: OkHttpClient): ApolloClient {
    val sqlCacheStore = SqlNormalizedCacheFactory("graphql.db")
    val baseApi = BuildConfig.apilocal
    val baseWssApi = BuildConfig.wssapilocal
    return ApolloClient.Builder().okHttpClient(okHttpClient = httpClient)
        .httpServerUrl("${baseApi}/graphql").webSocketServerUrl("${baseWssApi}/graphql")
        .wsProtocol(GraphQLWsProtocol.Factory()).webSocketReopenWhen { _, attempt ->
            CoroutineScope(Dispatchers.IO).launch {
                delay(attempt * 1000)
            }
            true
        }.normalizedCache(sqlCacheStore).build()
}