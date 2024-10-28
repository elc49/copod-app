package com.lomolo.copodv2.network

import com.apollographql.apollo3.ApolloClient
import com.apollographql.apollo3.api.ApolloResponse
import com.lomolo.copodv2.GetLocalLandsQuery

interface GraphQL {
    suspend fun getLocalLands(): ApolloResponse<GetLocalLandsQuery.Data>
}

class GraphQLServiceImpl(
    private val apolloClient: ApolloClient
) : GraphQL {
    override suspend fun getLocalLands() = apolloClient.query(GetLocalLandsQuery()).execute()
}