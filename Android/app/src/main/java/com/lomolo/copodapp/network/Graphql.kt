package com.lomolo.copodapp.network

import com.apollographql.apollo.ApolloClient
import com.apollographql.apollo.api.ApolloResponse
import com.apollographql.apollo.cache.normalized.FetchPolicy
import com.apollographql.apollo.cache.normalized.fetchPolicy
import com.lomolo.copodapp.ChargeMpesaMutation
import com.lomolo.copodapp.CreateOnboardingMutation
import com.lomolo.copodapp.GetIsTitleVerifiedQuery
import com.lomolo.copodapp.GetLandTitleDetailsQuery
import com.lomolo.copodapp.GetOnboardingByEmailAndVerificationQuery
import com.lomolo.copodapp.GetUserLandQuery
import com.lomolo.copodapp.PaymentUpdateSubscription
import com.lomolo.copodapp.UpdateDisplayPictureByIDMutation
import com.lomolo.copodapp.UpdateSupportingDocByIDMutation
import com.lomolo.copodapp.UpdateTitleDeedByIDMutation
import com.lomolo.copodapp.type.CreateOnboardingInput
import com.lomolo.copodapp.type.GetOnboardingByEmailAndVerificationInput
import com.lomolo.copodapp.type.GetUserLandsInput
import com.lomolo.copodapp.type.PayWithMpesaInput
import com.lomolo.copodapp.type.UpdateDisplayPictureByIDInput
import com.lomolo.copodapp.type.UpdateSupportingDocByIDInput
import com.lomolo.copodapp.type.UpdateTitleDeedByIDInput
import kotlinx.coroutines.flow.Flow

interface IGraphQL {
    suspend fun getUserLands(input: GetUserLandsInput): ApolloResponse<GetUserLandQuery.Data>
    suspend fun chargeMpesa(input: PayWithMpesaInput): ApolloResponse<ChargeMpesaMutation.Data>
    fun paymentUpdate(email: String): Flow<ApolloResponse<PaymentUpdateSubscription.Data>>
    suspend fun getOnboardingByEmailAndVerification(input: GetOnboardingByEmailAndVerificationInput): ApolloResponse<GetOnboardingByEmailAndVerificationQuery.Data>
    suspend fun createOnboarding(input: CreateOnboardingInput): ApolloResponse<CreateOnboardingMutation.Data>
    suspend fun getIsTitleVerified(titleNo: String): ApolloResponse<GetIsTitleVerifiedQuery.Data>
    suspend fun getLandTitleDetails(titleNo: String): ApolloResponse<GetLandTitleDetailsQuery.Data>
    suspend fun updateTitleDeedByID(input: UpdateTitleDeedByIDInput): ApolloResponse<UpdateTitleDeedByIDMutation.Data>
    suspend fun updateDisplayPictureByID(input: UpdateDisplayPictureByIDInput): ApolloResponse<UpdateDisplayPictureByIDMutation.Data>
    suspend fun updateSupportingDocByID(input: UpdateSupportingDocByIDInput): ApolloResponse<UpdateSupportingDocByIDMutation.Data>
}

class GraphQLServiceImpl(
    private val apolloClient: ApolloClient
) : IGraphQL {
    override suspend fun getUserLands(input: GetUserLandsInput) =
        apolloClient.query(GetUserLandQuery(input)).fetchPolicy(
            FetchPolicy.NetworkFirst
        ).execute()

    override suspend fun chargeMpesa(input: PayWithMpesaInput) =
        apolloClient.mutation(ChargeMpesaMutation(input)).execute()

    override fun paymentUpdate(email: String) = apolloClient.subscription(
        PaymentUpdateSubscription(email)
    ).toFlow()

    override suspend fun getOnboardingByEmailAndVerification(input: GetOnboardingByEmailAndVerificationInput) =
        apolloClient.query(GetOnboardingByEmailAndVerificationQuery(input)).fetchPolicy(
            FetchPolicy.NetworkFirst
        ).execute()

    override suspend fun createOnboarding(input: CreateOnboardingInput) =
        apolloClient.mutation(CreateOnboardingMutation(input)).execute()

    override suspend fun getIsTitleVerified(titleNo: String) =
        apolloClient.query(GetIsTitleVerifiedQuery(titleNo)).fetchPolicy(
            FetchPolicy.NetworkFirst
        ).execute()

    override suspend fun getLandTitleDetails(titleNo: String) =
        apolloClient.query(GetLandTitleDetailsQuery(titleNo)).fetchPolicy(
            FetchPolicy.NetworkFirst
        ).execute()

    override suspend fun updateTitleDeedByID(input: UpdateTitleDeedByIDInput) = apolloClient
        .mutation(UpdateTitleDeedByIDMutation(input))
        .execute()

    override suspend fun updateSupportingDocByID(input: UpdateSupportingDocByIDInput) = apolloClient
        .mutation(UpdateSupportingDocByIDMutation(input))
        .execute()

    override suspend fun updateDisplayPictureByID(input: UpdateDisplayPictureByIDInput) = apolloClient
        .mutation(UpdateDisplayPictureByIDMutation(input))
        .execute()
}