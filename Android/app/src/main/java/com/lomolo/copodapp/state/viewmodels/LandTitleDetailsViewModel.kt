package com.lomolo.copodapp.state.viewmodels

import android.util.Log
import androidx.compose.runtime.getValue
import androidx.compose.runtime.mutableStateOf
import androidx.compose.runtime.setValue
import androidx.lifecycle.ViewModel
import androidx.lifecycle.viewModelScope
import com.apollographql.apollo.exception.ApolloException
import com.lomolo.copodapp.GetLandTitleDetailsQuery
import com.lomolo.copodapp.network.IGraphQL
import kotlinx.coroutines.flow.MutableStateFlow
import kotlinx.coroutines.flow.StateFlow
import kotlinx.coroutines.flow.asStateFlow
import kotlinx.coroutines.flow.update
import kotlinx.coroutines.launch

interface GettingTitleDetails {
    data object Success: GettingTitleDetails
    data object Loading: GettingTitleDetails
    data class Error(val msg: String?): GettingTitleDetails
}

class LandTitleDetailsViewModel(
    private val graphqlApiService: IGraphQL,
) : ViewModel() {
    var gettingTitleDetails: GettingTitleDetails by mutableStateOf(GettingTitleDetails.Success)
        private set

    private val _titleDetails: MutableStateFlow<GetLandTitleDetailsQuery.GetLandTitleDetails?> =
        MutableStateFlow(null)
    val titleDetails: StateFlow<GetLandTitleDetailsQuery.GetLandTitleDetails?> = _titleDetails.asStateFlow()

    fun getLandTitleDetails(title: String) {
        if (gettingTitleDetails !is GettingTitleDetails.Loading) {
            gettingTitleDetails = GettingTitleDetails.Loading
            viewModelScope.launch {
                gettingTitleDetails = try {
                    val res = graphqlApiService.getLandTitleDetails(title).dataOrThrow()
                    _titleDetails.update { res.getLandTitleDetails }
                    GettingTitleDetails.Success
                } catch (e: ApolloException) {
                    e.printStackTrace()
                    Log.d(TAG, e.message ?: "Something went wrong")
                    GettingTitleDetails.Error(e.message)
                }
            }
        }
    }

    companion object {
        private const val TAG = "LandTitleDetailsViewModel"
    }
}