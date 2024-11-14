package com.lomolo.copodapp.ui.viewmodels

import android.util.Log
import androidx.compose.runtime.mutableStateOf
import androidx.compose.runtime.getValue
import androidx.compose.runtime.setValue
import androidx.lifecycle.SavedStateHandle
import androidx.lifecycle.ViewModel
import androidx.lifecycle.viewModelScope
import com.apollographql.apollo.exception.ApolloException
import com.lomolo.copodapp.model.DeviceDetails
import com.lomolo.copodapp.network.IGraphQL
import com.lomolo.copodapp.repository.IWeb3Auth
import com.lomolo.copodapp.type.PayWithMpesaInput
import com.lomolo.copodapp.type.PaymentReason
import com.lomolo.copodapp.ui.screens.MpesaScreenDestination
import com.lomolo.copodapp.util.Phone
import kotlinx.coroutines.flow.MutableStateFlow
import kotlinx.coroutines.flow.StateFlow
import kotlinx.coroutines.flow.asStateFlow
import kotlinx.coroutines.flow.update
import kotlinx.coroutines.launch

interface ChargingMpesa {
    data object Loading: ChargingMpesa
    data object Success: ChargingMpesa
    data object Paying: ChargingMpesa
    data object Paid: ChargingMpesa
    data class Error(val msg: String?): ChargingMpesa
}

data class Mpesa(
    val phone: String = "",
)

class MpesaViewModel(
    private val graphqlApiService: IGraphQL,
    private val web3Auth: IWeb3Auth,
    savedStateHandle: SavedStateHandle,
): ViewModel() {
    private val _mpesa: MutableStateFlow<Mpesa> = MutableStateFlow(Mpesa())
    val mpesa: StateFlow<Mpesa> = _mpesa.asStateFlow()

    var chargingMpesa: ChargingMpesa by mutableStateOf(ChargingMpesa.Success)
        private set

    private val uploadId: String = checkNotNull(
        savedStateHandle[MpesaScreenDestination.LAND_TITLE_ID_ARG]
    )

    fun setPhone(phone: String) {
        _mpesa.update { it.copy(phone = phone) }
    }

    fun isValidPhone(uiState: Mpesa, deviceDetails: DeviceDetails): Boolean {
        return with(uiState) {
            Phone.isValid(
                phone,
                deviceDetails.countryCode,
                deviceDetails.callingCode,
            )
        }
    }

    fun chargeMpesa(email: String, wallet: String, deviceDetails: DeviceDetails) {
        if (chargingMpesa !is ChargingMpesa.Loading && chargingMpesa !is ChargingMpesa.Paying) {
            chargingMpesa = ChargingMpesa.Loading
            viewModelScope.launch {
                chargingMpesa = try {
                    val phone = Phone.formatPhone(_mpesa.value.phone, deviceDetails.countryCode)
                    val input = PayWithMpesaInput(
                        reason = PaymentReason.LAND_REGISTRY,
                        phone = phone,
                        email = email,
                        walletAddress = wallet,
                        currency = deviceDetails.currency,
                        paymentFor = uploadId,
                    )
                    graphqlApiService.chargeMpesa(input)
                    ChargingMpesa.Paying
                } catch (e: ApolloException) {
                    Log.d(TAG, e.message ?: "Something went wrong")
                    ChargingMpesa.Error(e.message)
                }
            }
        }
    }

    companion object {
        private const val TAG = "MpesaViewModel"
    }

    init {
        viewModelScope.launch {
            try {
                val credentials = web3Auth.getCredentials(web3Auth.getPrivateKey())
                graphqlApiService.paymentUpdate(credentials.address).collect {
                    Log.d(TAG, it.data?.paymentUpdate.toString())
                    val data = it.data?.paymentUpdate
                    chargingMpesa = when(data?.status) {
                        "success" -> {
                            ChargingMpesa.Paid
                        }

                        else -> {
                            ChargingMpesa.Success
                        }
                    }
                }
            } catch (e: Exception) {
                Log.d(TAG, e.message ?: "Something went wrong")
            }
        }
    }
}