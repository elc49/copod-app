package com.lomolo.copodapp.state.viewmodels

import android.util.Log
import androidx.compose.runtime.getValue
import androidx.compose.runtime.mutableStateOf
import androidx.compose.runtime.setValue
import androidx.lifecycle.ViewModel
import androidx.lifecycle.viewModelScope
import com.apollographql.apollo.exception.ApolloException
import com.lomolo.copodapp.network.IGraphQL
import kotlinx.coroutines.flow.MutableStateFlow
import kotlinx.coroutines.flow.StateFlow
import kotlinx.coroutines.flow.asStateFlow
import kotlinx.coroutines.flow.update
import kotlinx.coroutines.launch

interface SearchingLand {
    data object Success: SearchingLand
    data object Loading: SearchingLand
    data class Error(val msg: String?): SearchingLand
}

class SearchLandViewModel(
    private val graphqlApiService: IGraphQL
): ViewModel() {
    private val _searchResult: MutableStateFlow<Boolean> = MutableStateFlow(false)
    val searchResult: StateFlow<Boolean> = _searchResult.asStateFlow()

    private val _searchQuery: MutableStateFlow<String> = MutableStateFlow("")
    val searchQuery: StateFlow<String> = _searchQuery.asStateFlow()

    var searchingLand: SearchingLand by mutableStateOf(SearchingLand.Success)
        private set

    fun updateSearchQuery(query: String) {
        if (searchingLand !is SearchingLand.Loading) _searchQuery.update { query }
    }

    fun searchLandTitle() {
        if (searchingLand !is SearchingLand.Loading && _searchQuery.value.isNotEmpty()) {
            searchingLand = SearchingLand.Loading
            viewModelScope.launch {
                searchingLand = try {
                    val res = graphqlApiService.getIsTitleVerified(_searchQuery.value.lowercase()).dataOrThrow()
                    _searchResult.update { res.getIsTitleVerified }
                    SearchingLand.Success
                } catch (e: ApolloException) {
                    e.printStackTrace()
                    Log.d(TAG, e.message ?: "Something went wrong")
                    SearchingLand.Error(e.message)
                }
            }
        }
    }

    fun resetState() {
        _searchQuery.value = ""
        _searchResult.value = false
    }

    companion object {
        private const val TAG = "SearchLandViewModel"
    }
}