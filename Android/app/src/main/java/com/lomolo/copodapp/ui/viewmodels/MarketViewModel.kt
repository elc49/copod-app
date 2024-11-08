package com.lomolo.copodapp.ui.viewmodels

import android.util.Log
import androidx.compose.runtime.getValue
import androidx.compose.runtime.mutableStateOf
import androidx.compose.runtime.setValue
import androidx.lifecycle.ViewModel
import androidx.lifecycle.viewModelScope
import com.apollographql.apollo.exception.ApolloException
import com.lomolo.copodapp.GetLocalLandsQuery
import com.lomolo.copodapp.network.IGraphQL
import kotlinx.coroutines.flow.MutableStateFlow
import kotlinx.coroutines.flow.StateFlow
import kotlinx.coroutines.flow.asStateFlow
import kotlinx.coroutines.flow.update
import kotlinx.coroutines.launch

interface GetLocalLands {
    data object Success: GetLocalLands
    data object Loading: GetLocalLands
    data class Error(val msg: String?): GetLocalLands
}

class MarketViewModel(
    private val graphqlService: IGraphQL
): ViewModel() {
    private val _lands: MutableStateFlow<List<GetLocalLandsQuery.GetLocalLand>> = MutableStateFlow(listOf())
    val lands: StateFlow<List<GetLocalLandsQuery.GetLocalLand>> = _lands.asStateFlow()

    var gettingLands: GetLocalLands by mutableStateOf(GetLocalLands.Success)
        private set

    fun getLocalLands() {
        if (gettingLands !is GetLocalLands.Loading) {
            gettingLands = GetLocalLands.Loading
            viewModelScope.launch {
                gettingLands = try {
                    val res = graphqlService.getLocalLands().dataOrThrow()
                    _lands.update { res.getLocalLands }
                    GetLocalLands.Success
                } catch (e: ApolloException) {
                    Log.d(TAG, e.message ?: "Something went wrong")
                    GetLocalLands.Error(e.message)
                }
            }
        }
    }

    companion object {
        private const val TAG = "MarketViewModel"
    }

    init {
        getLocalLands()
    }
}