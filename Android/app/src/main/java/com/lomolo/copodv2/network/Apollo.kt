package com.lomolo.copodv2.network

import com.apollographql.apollo3.ApolloClient

interface GraphQL {}

class GraphQLServiceImpl(
    private val apolloClient: ApolloClient
): GraphQL {}