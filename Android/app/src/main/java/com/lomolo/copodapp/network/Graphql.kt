package com.lomolo.copodapp.network

import com.apollographql.apollo3.ApolloClient
import com.apollographql.apollo3.api.ApolloResponse
import com.apollographql.apollo3.cache.normalized.FetchPolicy
import com.apollographql.apollo3.cache.normalized.fetchPolicy
import com.lomolo.copodapp.GetLocalLandsQuery
import com.lomolo.copodapp.GetUserLandQuery

interface IGraphQL {
    suspend fun getLocalLands(): ApolloResponse<GetLocalLandsQuery.Data>
    suspend fun getUserLands(email: String): ApolloResponse<GetUserLandQuery.Data>
}

class GraphQLServiceImpl(
    private val apolloClient: ApolloClient
) : IGraphQL {
    override suspend fun getLocalLands() = apolloClient.query(GetLocalLandsQuery()).fetchPolicy(
        FetchPolicy.NetworkFirst
    ).execute()

    override suspend fun getUserLands(email: String) = apolloClient.query(GetUserLandQuery(email)).fetchPolicy(
        FetchPolicy.NetworkFirst
    ).execute()
}