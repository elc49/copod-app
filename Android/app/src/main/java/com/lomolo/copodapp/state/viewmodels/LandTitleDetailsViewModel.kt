package com.lomolo.copodapp.state.viewmodels

import androidx.lifecycle.SavedStateHandle
import androidx.lifecycle.ViewModel
import com.lomolo.copodapp.ui.screens.FoundLandScreenDestination

class LandTitleDetailsViewModel(
    savedStateHandle: SavedStateHandle,
) : ViewModel() {
    val titleNo: String =
        checkNotNull(savedStateHandle[FoundLandScreenDestination.TITLE_NO_ARG])
}