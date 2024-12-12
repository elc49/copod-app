package com.lomolo.copodapp.state.viewmodels

import android.util.Log
import androidx.compose.runtime.getValue
import androidx.compose.runtime.mutableStateOf
import androidx.compose.runtime.setValue
import androidx.lifecycle.ViewModel
import androidx.lifecycle.viewModelScope
import com.apollographql.apollo.exception.ApolloException
import com.lomolo.copodapp.GetOnboardingByEmailQuery
import com.lomolo.copodapp.network.IGraphQL
import com.lomolo.copodapp.network.IRestFul
import com.lomolo.copodapp.repository.IWeb3Auth
import com.lomolo.copodapp.type.CreateOnboardingInput
import com.web3auth.core.types.UserInfo
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

interface GetCurrentOnboarding {
    data object Success: GetCurrentOnboarding
    data object Loading: GetCurrentOnboarding
    data class Error(val msg: String?): GetCurrentOnboarding
}

interface Onboarding {
    data object Success: Onboarding
    data object Loading: Onboarding
    data class Error(val msg: String?): Onboarding
}

class OnboardingViewModel(
    private val restApiService: IRestFul,
    private val graphqlApiService: IGraphQL,
    web3Auth: IWeb3Auth,
) : ViewModel() {
    private val userInfo: UserInfo = web3Auth.getUserInfo()

    private val _landTitle: MutableStateFlow<String> = MutableStateFlow("")
    val landTitle: StateFlow<String> = _landTitle.asStateFlow()

    private val _supportingDoc: MutableStateFlow<String> = MutableStateFlow("")
    val supportingDoc: StateFlow<String> = _supportingDoc.asStateFlow()

    private val _displayPicture: MutableStateFlow<String> = MutableStateFlow("")
    val displayPicture: StateFlow<String> = _displayPicture.asStateFlow()

    private val _currentOnboarding: MutableStateFlow<GetOnboardingByEmailQuery.GetOnboardingByEmail?> = MutableStateFlow(null)
    val currentOnboarding: StateFlow<GetOnboardingByEmailQuery.GetOnboardingByEmail?> = _currentOnboarding.asStateFlow()

    var uploadingLandDoc: UploadingDoc by mutableStateOf(UploadingDoc.Success)
        private set

    var uploadingGovtId: UploadingDoc by mutableStateOf(UploadingDoc.Success)
        private set

    var uploadingDp: UploadingDoc by mutableStateOf(UploadingDoc.Success)
        private set

    var gettingCurrentOnboarding: GetCurrentOnboarding by mutableStateOf(GetCurrentOnboarding.Success)
        private set

    var onboarding: Onboarding by mutableStateOf(Onboarding.Success)
        private set

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
                    _landTitle.emit(res.imageUri)
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
                    _supportingDoc.emit(res.imageUri)
                    UploadingDoc.Success
                } catch (e: Exception) {
                    Log.d(TAG, e.message ?: "Something went wrong")
                    UploadingDoc.Error(e.message ?: "Something went wrong")
                }
            }
        }
    }

    fun uploadDisplayPicture(fileName: String, stream: InputStream) {
        if (uploadingDp !is UploadingDoc.Loading) {
            uploadingDp = UploadingDoc.Loading
            val request = stream.readBytes().toRequestBody()
            val filePart = MultipartBody.Part.createFormData(
                "file",
                "$fileName}.jpg",
                request,
            )
            viewModelScope.launch {
                uploadingDp = try {
                    val res = restApiService.uploadDoc(filePart)
                    _displayPicture.emit(res.imageUri)
                    UploadingDoc.Success
                } catch (e: Exception) {
                    Log.d(TAG, e.message ?: "Something went wrong")
                    UploadingDoc.Error(e.message ?: "Something went wrong")
                }
            }
        }
    }

    fun getCurrentOnboarding() {
        if (gettingCurrentOnboarding !is GetCurrentOnboarding.Loading) {
            gettingCurrentOnboarding = GetCurrentOnboarding.Loading
            viewModelScope.launch {
                gettingCurrentOnboarding = try {
                    val res = graphqlApiService.getOnboardingByEmail(userInfo.email).dataOrThrow()
                    _currentOnboarding.emit(res.getOnboardingByEmail)
                    GetCurrentOnboarding.Success
                } catch (e: ApolloException) {
                    Log.d(TAG, e.message ?: "Something went wrong")
                    GetCurrentOnboarding.Error(e.message)
                }
            }
        }
    }

    fun createOnboarding(cb: () -> Unit = {}) {
        if (onboarding !is Onboarding.Loading) {
            onboarding = Onboarding.Loading
            viewModelScope.launch {
                onboarding = try {
                    graphqlApiService.createOnboarding(
                        CreateOnboardingInput(
                            email = userInfo.email,
                            titleUrl = _landTitle.value,
                            displayPictureUrl = _displayPicture.value,
                            supportdocUrl = _supportingDoc.value
                        )
                    )
                    Onboarding.Success.also { cb() }
                } catch (e: ApolloException) {
                    Log.d(TAG, e.message ?: "Something went wrong")
                    Onboarding.Error(e.message)
                }
            }
        }
    }

    companion object {
        private const val TAG = "OnboardingViewModel"
    }
}