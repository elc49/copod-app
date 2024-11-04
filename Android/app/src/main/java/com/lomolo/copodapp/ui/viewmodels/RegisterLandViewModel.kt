package com.lomolo.copodapp.ui.viewmodels

import android.util.Log
import androidx.compose.runtime.getValue
import androidx.compose.runtime.mutableStateOf
import androidx.compose.runtime.setValue
import androidx.lifecycle.ViewModel
import androidx.lifecycle.viewModelScope
import com.lomolo.copodapp.network.IRestFul
import kotlinx.coroutines.flow.MutableStateFlow
import kotlinx.coroutines.flow.StateFlow
import kotlinx.coroutines.flow.asStateFlow
import kotlinx.coroutines.launch
import okhttp3.MultipartBody
import okhttp3.RequestBody.Companion.toRequestBody
import java.io.InputStream
import java.lang.Exception

interface UploadingDoc {
    data object Success : UploadingDoc
    data object Loading : UploadingDoc
    data class Error(val msg: String?) : UploadingDoc
}

data class UploadDocState(
    val state: Map<String, UploadDocState> = mapOf(),
)

class RegisterLandViewModel(
    private val IRestFul: IRestFul,
) : ViewModel() {
    private val _uploadState: MutableStateFlow<UploadDocState> = MutableStateFlow(UploadDocState())
    val uploadState: StateFlow<UploadDocState> = _uploadState.asStateFlow()

    var uploadingDoc: UploadingDoc by mutableStateOf(UploadingDoc.Success)
        private set
    var uploadingLandDoc: UploadingDoc by mutableStateOf(UploadingDoc.Success)
        private set
    var uploadingGovtId: UploadingDoc by mutableStateOf(UploadingDoc.Success)
        private set

    private fun uploadDoc(stream: InputStream) {
        val request = stream.readBytes().toRequestBody()
        val filePart = MultipartBody.Part.createFormData(
            "file",
            "${System.currentTimeMillis()}.jpg",
            request,
        )
        viewModelScope.launch {
            try {
                IRestFul.uploadDoc(filePart)
            } catch (e: Exception) {
                Log.d(TAG, e.message ?: "Something went wrong")
            }
        }
    }

    fun uploadLandTitle(stream: InputStream) {
        if (uploadingLandDoc !is UploadingDoc.Loading) {
            uploadingLandDoc = UploadingDoc.Loading
            viewModelScope.launch {
                uploadingLandDoc = try {
                    uploadDoc(stream)
                    UploadingDoc.Success
                } catch (e: Exception) {
                    Log.d(TAG, e.message ?: "Something went wrong")
                    UploadingDoc.Error(e.message)
                }
            }
        }
    }

    fun uploadGovtIssuedId(stream: InputStream) {
        if (uploadingGovtId !is UploadingDoc.Loading) {
            uploadingGovtId = UploadingDoc.Loading
            viewModelScope.launch {
                uploadingGovtId = try {
                    uploadDoc(stream)
                    UploadingDoc.Success
                } catch (e: Exception) {
                    Log.d(TAG, e.message ?: "Something went wrong")
                    UploadingDoc.Error(e.message)
                }
            }
        }
    }

    companion object {
        private const val TAG = "RegisterLandViewModel"
    }
}