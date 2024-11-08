package com.lomolo.copodapp.ui.viewmodels

import android.util.Log
import androidx.compose.runtime.getValue
import androidx.compose.runtime.mutableStateOf
import androidx.compose.runtime.setValue
import androidx.lifecycle.ViewModel
import androidx.lifecycle.viewModelScope
import com.apollographql.apollo.exception.ApolloException
import com.lomolo.copodapp.GetUserLandQuery
import com.lomolo.copodapp.network.IGraphQL
import kotlinx.coroutines.flow.MutableStateFlow
import kotlinx.coroutines.flow.StateFlow
import kotlinx.coroutines.flow.asStateFlow
import kotlinx.coroutines.launch

interface GetUserLands {
    data object Success: GetUserLands
    data object Loading: GetUserLands
    data class Error(val msg: String?): GetUserLands
}

class LandViewModel(
    private val graphqlService: IGraphQL,
) : ViewModel() {
    private val _lands: MutableStateFlow<List<GetUserLandQuery.GetUserLand>> = MutableStateFlow(listOf())
    val lands: StateFlow<List<GetUserLandQuery.GetUserLand>> = _lands.asStateFlow()

    var gettingUserLands: GetUserLands by mutableStateOf(GetUserLands.Success)
        private set

    fun getUserLands(email: String) {
        if (gettingUserLands !is GetUserLands.Loading) {
            gettingUserLands = GetUserLands.Loading
            viewModelScope.launch {
                gettingUserLands = try {
                    graphqlService.getUserLands(email)
                    GetUserLands.Success
                } catch (e: ApolloException) {
                    Log.d(TAG, e.message ?: "Something went wrong")
                    GetUserLands.Error(e.message)
                }
            }
        }
    }

    companion object {
        private const val TAG = "LandViewModel"
    }
}