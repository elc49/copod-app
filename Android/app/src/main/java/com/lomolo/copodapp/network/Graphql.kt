package com.lomolo.copodapp.network

import com.apollographql.apollo.ApolloClient
import com.apollographql.apollo.api.ApolloResponse
import com.apollographql.apollo.cache.normalized.FetchPolicy
import com.apollographql.apollo.cache.normalized.fetchPolicy
import com.lomolo.copodapp.ChargeMpesaMutation
import com.lomolo.copodapp.GetUserLandQuery
import com.lomolo.copodapp.PaymentUpdateSubscription
import com.lomolo.copodapp.type.PayWithMpesaInput
import kotlinx.coroutines.flow.Flow

interface IGraphQL {
    suspend fun getUserLands(email: String): ApolloResponse<GetUserLandQuery.Data>
    suspend fun chargeMpesa(input: PayWithMpesaInput): ApolloResponse<ChargeMpesaMutation.Data>
    fun paymentUpdate(email: String): Flow<ApolloResponse<PaymentUpdateSubscription.Data>>
}

class GraphQLServiceImpl(
    private val apolloClient: ApolloClient
) : IGraphQL {
    override suspend fun getUserLands(email: String) =
        apolloClient.query(GetUserLandQuery(email)).fetchPolicy(
            FetchPolicy.NetworkFirst
        ).execute()

    override suspend fun chargeMpesa(input: PayWithMpesaInput) =
        apolloClient.mutation(ChargeMpesaMutation(input)).execute()

    override fun paymentUpdate(email: String) = apolloClient.subscription(
        PaymentUpdateSubscription(email)
    ).toFlow()
}