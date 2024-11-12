package com.lomolo.copodapp.ui.viewmodels

import android.util.Log
import androidx.compose.runtime.getValue
import androidx.compose.runtime.mutableStateOf
import androidx.compose.runtime.setValue
import androidx.lifecycle.ViewModel
import androidx.lifecycle.viewModelScope
import com.apollographql.apollo.exception.ApolloException
import com.lomolo.copodapp.network.IGraphQL
import com.lomolo.copodapp.network.IRestFul
import com.lomolo.copodapp.repository.IWeb3Auth
import com.lomolo.copodapp.type.Doc
import com.lomolo.copodapp.type.UploadInput
import com.web3auth.core.types.UserInfo
import kotlinx.coroutines.flow.MutableStateFlow
import kotlinx.coroutines.flow.StateFlow
import kotlinx.coroutines.flow.asStateFlow
import kotlinx.coroutines.flow.update
import kotlinx.coroutines.launch
import okhttp3.MultipartBody
import okhttp3.RequestBody.Companion.toRequestBody
import okhttp3.internal.toImmutableMap
import java.io.InputStream
import java.lang.Exception

interface UploadingDoc {
    data object Success : UploadingDoc
    data object Loading : UploadingDoc
    data class Error(val msg: String?) : UploadingDoc
}

interface SaveUploads {
    data object Success: SaveUploads
    data object Loading: SaveUploads
    data class Error(val msg: String?): SaveUploads
}

data class UploadDocState(
    val images: Map<String, String> = mapOf(),
)

class RegisterLandViewModel(
    private val restApiService: IRestFul,
    private val graphqlApiService: IGraphQL,
    web3Auth: IWeb3Auth,
) : ViewModel() {
    private val _images: MutableStateFlow<UploadDocState> = MutableStateFlow(UploadDocState())
    val images: StateFlow<UploadDocState> = _images.asStateFlow()

    var uploadingLandDoc: UploadingDoc by mutableStateOf(UploadingDoc.Success)
        private set

    var uploadingGovtId: UploadingDoc by mutableStateOf(UploadingDoc.Success)
        private set

    var savingUploads: SaveUploads by mutableStateOf(SaveUploads.Success)
        private set

    private val userInfo: UserInfo = web3Auth.getUserInfo()

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
                    _images.update {
                        val m = it.images.toMutableMap()
                        m[Doc.LAND_TITLE.toString()] = res.imageUri
                        it.copy(images = m.toImmutableMap())
                    }
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
                    _images.update {
                        val m = it.images.toMutableMap()
                        m[Doc.GOVT_ID.toString()] = res.imageUri
                        it.copy(images = m.toImmutableMap())
                    }
                    UploadingDoc.Success
                } catch (e: Exception) {
                    Log.d(TAG, e.message ?: "Something went wrong")
                    UploadingDoc.Error(e.message ?: "Something went wrong")
                }
            }
        }
    }

    fun saveUploads(cb: (String) -> Unit) {
        if (savingUploads !is SaveUploads.Loading) {
           savingUploads = SaveUploads.Loading
            viewModelScope.launch {
                savingUploads = try {
                    val images = _images.value
                    val res = graphqlApiService.createUpload(
                        UploadInput(
                            type = Doc.LAND_TITLE,
                            images.images[Doc.LAND_TITLE.toString()]!!,
                            images.images[Doc.GOVT_ID.toString()]!!,
                            userInfo.email,
                        )
                    ).dataOrThrow()
                    SaveUploads.Success.also {
                        cb(res.createUpload.id.toString())
                    }
                } catch (e: ApolloException) {
                    Log.d(TAG, e.message ?: "Something went wrong")
                    SaveUploads.Error(e.message)
                }
            }
        }
    }

    companion object {
        private const val TAG = "RegisterLandViewModel"
    }
}