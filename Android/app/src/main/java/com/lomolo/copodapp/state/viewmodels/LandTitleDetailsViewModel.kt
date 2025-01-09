package com.lomolo.copodapp.state.viewmodels

import android.util.Log
import androidx.compose.runtime.mutableStateOf
import androidx.compose.runtime.setValue
import androidx.compose.runtime.getValue
import androidx.lifecycle.SavedStateHandle
import androidx.lifecycle.ViewModel
import androidx.lifecycle.viewModelScope
import com.apollographql.apollo.exception.ApolloException
import com.lomolo.copodapp.network.IGraphQL
import kotlinx.coroutines.launch

interface GettingTitleDetails {
    data object Success: GettingTitleDetails
    data object Loading: GettingTitleDetails
    data class Error(val msg: String?): GettingTitleDetails
}

class LandTitleDetailsViewModel(
    private val graphqlApiService: IGraphQL,
    savedStateHandle: SavedStateHandle,
) : ViewModel() {
    var gettingTitleDetails: GettingTitleDetails by mutableStateOf(GettingTitleDetails.Success)
        private set

    fun getLandTitleDetails() {
        if (gettingTitleDetails !is GettingTitleDetails.Loading) {
            gettingTitleDetails = GettingTitleDetails.Loading
            viewModelScope.launch {
                try {
                } catch (e: ApolloException) {
                    e.printStackTrace()
                    Log.d(TAG, e.message ?: "Something went wrong")
                }
            }
        }
    }

    companion object {
        private const val TAG = "LandTitleDetailsViewModel"
    }
}