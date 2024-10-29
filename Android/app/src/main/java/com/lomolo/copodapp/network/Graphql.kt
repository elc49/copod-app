package com.lomolo.copodapp.network

import com.apollographql.apollo3.ApolloClient
import com.apollographql.apollo3.api.ApolloResponse
import com.apollographql.apollo3.cache.normalized.FetchPolicy
import com.apollographql.apollo3.cache.normalized.fetchPolicy
import com.lomolo.copodapp.GetLocalLandsQuery

interface GraphQL {
    suspend fun getLocalLands(): ApolloResponse<GetLocalLandsQuery.Data>
}

class GraphQLServiceImpl(
    private val apolloClient: ApolloClient
) : GraphQL {
    override suspend fun getLocalLands() = apolloClient.query(GetLocalLandsQuery()).fetchPolicy(
        FetchPolicy.NetworkFirst
    ).execute()
}