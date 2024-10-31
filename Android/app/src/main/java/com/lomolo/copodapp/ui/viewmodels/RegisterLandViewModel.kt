package com.lomolo.copodapp.ui.viewmodels

import androidx.compose.runtime.getValue
import androidx.compose.runtime.mutableStateOf
import androidx.compose.runtime.setValue
import androidx.lifecycle.ViewModel
import com.lomolo.copodapp.network.RestFul
import kotlinx.coroutines.flow.MutableStateFlow
import kotlinx.coroutines.flow.StateFlow
import kotlinx.coroutines.flow.asStateFlow
import okhttp3.MultipartBody
import okhttp3.RequestBody.Companion.toRequestBody
import java.io.InputStream

interface UploadingDoc {
    data object Success : UploadingDoc
    data object Loading : UploadingDoc
    data class Error(val msg: String?) : UploadingDoc
}

data class UploadDocState(
    val state: Map<String, UploadDocState> = mapOf(),
)

class RegisterLandViewModel(
    private val restFul: RestFul,
) : ViewModel() {
    private val _uploadState: MutableStateFlow<UploadDocState> = MutableStateFlow(UploadDocState())
    val uploadState: StateFlow<UploadDocState> = _uploadState.asStateFlow()

    var uploadingDoc: UploadingDoc by mutableStateOf(UploadingDoc.Success)
        private set

    fun uploadDoc(stream: InputStream) {
        val request = stream.readBytes().toRequestBody()
        val filePart = MultipartBody.Part.createFormData(
            "file",
            "${System.currentTimeMillis()}.jpg",
            request,
        )
    }
}