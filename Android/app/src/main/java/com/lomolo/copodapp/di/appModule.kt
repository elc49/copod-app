package com.lomolo.copodapp.di

import com.apollographql.apollo.ApolloClient
import com.lomolo.copodapp.graphql.getGraphqlClient
import com.lomolo.copodapp.http.getHttpClient
import com.lomolo.copodapp.json.getJsonAdapter
import com.lomolo.copodapp.network.GraphQLServiceImpl
import com.lomolo.copodapp.network.IGraphQL
import com.lomolo.copodapp.network.IRestFul
import com.lomolo.copodapp.repository.IWeb3Auth
import com.lomolo.copodapp.rest.getRestService
import com.lomolo.copodapp.retrofit.getRestApiClient
import com.lomolo.copodapp.state.viewmodels.LandTitleDetailsViewModel
import com.lomolo.copodapp.state.viewmodels.LandViewModel
import com.lomolo.copodapp.state.viewmodels.MainViewModel
import com.lomolo.copodapp.state.viewmodels.MarketViewModel
import com.lomolo.copodapp.state.viewmodels.MpesaViewModel
import com.lomolo.copodapp.state.viewmodels.OnboardingViewModel
import com.lomolo.copodapp.state.viewmodels.SearchLandViewModel
import com.lomolo.copodapp.web3.getWeb3AuthImpl
import com.squareup.moshi.Moshi
import okhttp3.OkHttpClient
import org.koin.core.module.dsl.viewModel
import org.koin.dsl.module
import retrofit2.Retrofit

val appModule = module {
    single<OkHttpClient> { getHttpClient() }
    single<ApolloClient> { getGraphqlClient(get()) }
    single<IWeb3Auth> { getWeb3AuthImpl(get()) }
    single<IGraphQL> { GraphQLServiceImpl(get()) }
    single<Retrofit> { getRestApiClient(get(), get()) }
    single<IRestFul> { getRestService(get()) }
    single<Moshi> { getJsonAdapter() }

    viewModel { MainViewModel(get(), get(), get()) }
    viewModel { MarketViewModel() }
    viewModel { LandViewModel(get()) }
    viewModel { MpesaViewModel(get(), get(), get()) }
    viewModel { OnboardingViewModel(get(), get(), get()) }
    viewModel { SearchLandViewModel(get()) }
    viewModel { LandTitleDetailsViewModel(get()) }
}