package com.lomolo.copodv2.di

import android.content.Context
import android.net.Uri
import com.apollographql.apollo3.ApolloClient
import com.apollographql.apollo3.cache.normalized.normalizedCache
import com.apollographql.apollo3.cache.normalized.sql.SqlNormalizedCacheFactory
import com.apollographql.apollo3.network.okHttpClient
import com.lomolo.copodv2.BuildConfig
import com.lomolo.copodv2.R
import com.lomolo.copodv2.network.GraphQL
import com.lomolo.copodv2.network.GraphQLServiceImpl
import com.lomolo.copodv2.repository.IWeb3Auth
import com.lomolo.copodv2.viewmodels.MainViewModel
import com.web3auth.core.Web3Auth
import com.web3auth.core.types.Network
import com.web3auth.core.types.Web3AuthOptions
import okhttp3.OkHttpClient
import org.koin.core.module.dsl.viewModel
import org.koin.dsl.module
import java.util.concurrent.TimeUnit

val appModule = module {
    single {
        getOkHttpClient()
    }
    single {
        getGraphqlService(get())
    }
    single {
        getWeb3AuthImpl(get())
    }
    single {
        getApolloClient(get())
    }

    viewModel { MainViewModel(get()) }
}

private fun getWeb3AuthImpl(context: Context): IWeb3Auth {
    val web3auth = Web3Auth(
        Web3AuthOptions(
            context = context,
            clientId = context.getString(R.string.web3auth_project_id),
            network = Network.SAPPHIRE_DEVNET,
            redirectUrl = Uri.parse("com.lomolo.copodv2://auth"),
        )
    )

    return com.lomolo.copodv2.repository.Web3Auth(web3auth)
}

private fun getGraphqlService(apolloClient: ApolloClient): GraphQL {
    return GraphQLServiceImpl(apolloClient)
}

private fun getOkHttpClient(): OkHttpClient {
    return OkHttpClient.Builder()
        .connectTimeout(5, TimeUnit.MINUTES)
        .callTimeout(5, TimeUnit.MINUTES)
        .readTimeout(5, TimeUnit.MINUTES)
        .writeTimeout(5, TimeUnit.MINUTES)
        .build()
}

private fun getApolloClient(httpClient: OkHttpClient): ApolloClient {
    val sqlCacheStore = SqlNormalizedCacheFactory("graphql.db")
    val baseApi = BuildConfig.apilocal
    return ApolloClient.Builder()
        .okHttpClient(okHttpClient = httpClient)
        .httpServerUrl("${baseApi}/graphql")
        .normalizedCache(sqlCacheStore)
        .build()
}