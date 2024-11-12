package com.lomolo.copodapp.network

import com.apollographql.apollo.ApolloClient
import com.apollographql.apollo.api.ApolloResponse
import com.apollographql.apollo.cache.normalized.FetchPolicy
import com.apollographql.apollo.cache.normalized.fetchPolicy
import com.lomolo.copodapp.ChargeMpesaMutation
import com.lomolo.copodapp.GetLocalLandsQuery
import com.lomolo.copodapp.GetUserLandQuery
import com.lomolo.copodapp.PaymentUpdateSubscription
import com.lomolo.copodapp.UploadLandTitleMutation
import com.lomolo.copodapp.UploadSupportingDocMutation
import com.lomolo.copodapp.type.DocUploadInput
import com.lomolo.copodapp.type.PayWithMpesaInput
import kotlinx.coroutines.flow.Flow

interface IGraphQL {
    suspend fun getLocalLands(): ApolloResponse<GetLocalLandsQuery.Data>
    suspend fun getUserLands(walletAddress: String): ApolloResponse<GetUserLandQuery.Data>
    suspend fun chargeMpesa(input: PayWithMpesaInput): ApolloResponse<ChargeMpesaMutation.Data>
    fun paymentUpdate(walletAddress: String): Flow<ApolloResponse<PaymentUpdateSubscription.Data>>
    suspend fun uploadLandTitle(input: DocUploadInput): ApolloResponse<UploadLandTitleMutation.Data>
    suspend fun uploadSupportingDoc(input: DocUploadInput): ApolloResponse<UploadSupportingDocMutation.Data>
}

class GraphQLServiceImpl(
    private val apolloClient: ApolloClient
) : IGraphQL {
    override suspend fun getLocalLands() = apolloClient.query(GetLocalLandsQuery()).fetchPolicy(
        FetchPolicy.NetworkFirst
    ).execute()

    override suspend fun getUserLands(walletAddress: String) =
        apolloClient.query(GetUserLandQuery(walletAddress)).fetchPolicy(
            FetchPolicy.NetworkFirst
        ).execute()

    override suspend fun chargeMpesa(input: PayWithMpesaInput) =
        apolloClient.mutation(ChargeMpesaMutation(input)).execute()

    override fun paymentUpdate(walletAddress: String) = apolloClient.subscription(
        PaymentUpdateSubscription(walletAddress)
    ).toFlow()

    override suspend fun uploadLandTitle(input: DocUploadInput) = apolloClient.mutation(
        UploadLandTitleMutation(input)
    ).execute()

    override suspend fun uploadSupportingDoc(input: DocUploadInput) = apolloClient.mutation(
        UploadSupportingDocMutation(input)
    ).execute()
}