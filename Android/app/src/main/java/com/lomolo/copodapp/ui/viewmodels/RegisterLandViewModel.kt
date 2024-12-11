package com.lomolo.copodapp.ui.viewmodels

import android.util.Log
import androidx.compose.runtime.getValue
import androidx.compose.runtime.mutableStateOf
import androidx.compose.runtime.setValue
import androidx.lifecycle.SavedStateHandle
import androidx.lifecycle.ViewModel
import androidx.lifecycle.viewModelScope
import com.lomolo.copodapp.network.IGraphQL
import com.lomolo.copodapp.network.IRestFul
import com.lomolo.copodapp.ui.screens.UploadGovtIssuedIdScreenDestination
import kotlinx.coroutines.flow.MutableStateFlow
import kotlinx.coroutines.flow.StateFlow
import kotlinx.coroutines.flow.asStateFlow
import kotlinx.coroutines.flow.update
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

interface SaveUpload {
    data object Success : SaveUpload
    data object Loading : SaveUpload
    data class Error(val msg: String?) : SaveUpload
}

class RegisterLandViewModel(
    private val restApiService: IRestFul,
    private val graphqlApiService: IGraphQL,
    savedStateHandle: SavedStateHandle,
) : ViewModel() {
    private val _landTitle: MutableStateFlow<String> = MutableStateFlow("")
    val landTitle: StateFlow<String> = _landTitle.asStateFlow()

    private val _supportingDoc: MutableStateFlow<String> = MutableStateFlow("")
    val supportingDoc: StateFlow<String> = _supportingDoc.asStateFlow()

    var uploadingLandDoc: UploadingDoc by mutableStateOf(UploadingDoc.Success)
        private set

    var uploadingGovtId: UploadingDoc by mutableStateOf(UploadingDoc.Success)
        private set

    var savingLandTitle: SaveUpload by mutableStateOf(SaveUpload.Success)
        private set

    var savingSupportingDoc: SaveUpload by mutableStateOf(SaveUpload.Success)
        private set

    private val titleId: String? = savedStateHandle[UploadGovtIssuedIdScreenDestination.LAND_TITLE_ID_ARG]

    fun uploadLandTitle(fileName: String, stream: InputStream) {
        if (uploadingLandDoc !is UploadingDoc.Loading) {
            uploadingLandDoc = UploadingDoc.Loading
            val request = stream.readBytes().toRequestBody()
            val filePart = MultipartBody.Part.createFormData(
                "file",
                "${fileName}.jpg",
                request,
            )
            viewModelScope.launch {
                uploadingLandDoc = try {
                    val res = restApiService.uploadDoc(filePart)
                    _landTitle.update { res.imageUri }
                    UploadingDoc.Success
                } catch (e: Exception) {
                    Log.d(TAG, e.message ?: "Something went wrong")
                    UploadingDoc.Error(e.message ?: "Something went wrong")
                }
            }
        }
    }

    fun uploadGovtIssuedId(fileName: String, stream: InputStream) {
        if (uploadingGovtId !is UploadingDoc.Loading) {
            uploadingGovtId = UploadingDoc.Loading
            val request = stream.readBytes().toRequestBody()
            val filePart = MultipartBody.Part.createFormData(
                "file",
                "$fileName}.jpg",
                request,
            )
            viewModelScope.launch {
                uploadingGovtId = try {
                    val res = restApiService.uploadDoc(filePart)
                    _supportingDoc.update { res.imageUri }
                    UploadingDoc.Success
                } catch (e: Exception) {
                    Log.d(TAG, e.message ?: "Something went wrong")
                    UploadingDoc.Error(e.message ?: "Something went wrong")
                }
            }
        }
    }

    companion object {
        private const val TAG = "RegisterLandViewModel"
    }
}