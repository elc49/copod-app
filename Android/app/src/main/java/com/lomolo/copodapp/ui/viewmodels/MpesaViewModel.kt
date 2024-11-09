package com.lomolo.copodapp.ui.viewmodels

import android.util.Log
import androidx.compose.runtime.mutableStateOf
import androidx.compose.runtime.getValue
import androidx.compose.runtime.setValue
import androidx.lifecycle.ViewModel
import androidx.lifecycle.viewModelScope
import com.apollographql.apollo.exception.ApolloException
import com.lomolo.copodapp.network.IGraphQL
import com.lomolo.copodapp.type.PayWithMpesaInput
import com.lomolo.copodapp.type.PaymentReason
import kotlinx.coroutines.flow.MutableStateFlow
import kotlinx.coroutines.flow.StateFlow
import kotlinx.coroutines.flow.asStateFlow
import kotlinx.coroutines.flow.update
import kotlinx.coroutines.launch

interface ChargingMpesa {
    data object Loading: ChargingMpesa
    data object Success: ChargingMpesa
    data class Error(val msg: String?): ChargingMpesa
}

data class Mpesa(
    val phone: String = "",
)

class MpesaViewModel(
    private val graphqlApiService: IGraphQL,
): ViewModel() {
    private val _mpesa: MutableStateFlow<Mpesa> = MutableStateFlow(Mpesa())
    val mpesa: StateFlow<Mpesa> = _mpesa.asStateFlow()

    var chargingMpesa: ChargingMpesa by mutableStateOf(ChargingMpesa.Success)
        private set

    fun setPhone(phone: String) {
        _mpesa.update { it.copy(phone = phone) }
    }

    fun chargeMpesa() {
        if (chargingMpesa !is ChargingMpesa.Loading) {
            chargingMpesa = ChargingMpesa.Loading
            viewModelScope.launch {
                chargingMpesa = try {
                    val input = PayWithMpesaInput(
                        reason = PaymentReason.LAND_REGISTRATION,
                        phone = _mpesa.value.phone,
                        amount = 0,
                        email = "",
                        currency = "KES",
                    )
                    graphqlApiService.chargeMpesa(input)
                    ChargingMpesa.Success
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
}