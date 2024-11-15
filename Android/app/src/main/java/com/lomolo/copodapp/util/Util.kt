package com.lomolo.copodapp.util

import android.icu.number.Notation
import android.icu.number.NumberFormatter
import android.icu.number.Precision
import android.os.Build
import android.util.Log
import com.google.i18n.phonenumbers.PhoneNumberUtil
import com.google.i18n.phonenumbers.PhoneNumberUtil.PhoneNumberFormat
import com.google.i18n.phonenumbers.Phonenumber.PhoneNumber
import java.text.NumberFormat
import java.util.Currency
import java.util.Locale

object Util {
    fun capitalize(text: String): String {
        return text
            .lowercase()
            .replaceFirstChar {
                if (it.isLowerCase())
                    it.titlecase(Locale.getDefault())
                else
                    it.toString()
            }
    }

    fun formatCurrency(currency: String, amount: Int): String {
        val numberFormat = NumberFormat.getCurrencyInstance(Locale.getDefault())
        numberFormat.maximumFractionDigits = 0
        numberFormat.currency = Currency.getInstance(currency)

        return if (Build.VERSION.SDK_INT >= Build.VERSION_CODES.R) {
            NumberFormatter.with().notation(Notation.simple())
                .unit(android.icu.util.Currency.getInstance(currency))
                .precision(Precision.maxFraction(0)).locale(Locale.US).format(amount).toString()
        } else {
            numberFormat.format(amount).toString()
        }
    }


    object Phone {
        val phoneUtil = PhoneNumberUtil.getInstance()
        private const val TAG = "Phone"

        private fun parsePhoneNumber(phone: String, countryCode: String): PhoneNumber {
            var n = PhoneNumber()

            try {
                n = phoneUtil.parse(phone, countryCode)
            } catch (e: Exception) {
                Log.d(TAG, e.message ?: "Something went wrong")
            }
            return n
        }

        fun isValid(phone: String, countryCode: String, callingCode: String): Boolean {
            return try {
                if (phone.isEmpty()) return false
                val p = PhoneNumber()
                p.countryCode = callingCode.filter { it.toString() != "+" }.toInt()
                p.nationalNumber = phone.toLong()
                return phoneUtil.isValidNumber(parsePhoneNumber(phone, countryCode))
            } catch (e: Exception) {
                Log.d(TAG, e.message ?: "Something went wrong")
                false
            }
        }

        fun formatPhone(phone: String, countryCode: String): String {
            var p: PhoneNumber = PhoneNumber()
            return try {
                p = parsePhoneNumber(phone, countryCode)
                phoneUtil.format(parsePhoneNumber(phone, countryCode), PhoneNumberFormat.E164)
            } catch (e: Exception) {
                Log.d(TAG, e.message ?: "Something went wrong")
                "+"+p.countryCode.toString()+p.nationalNumber.toString()
            }
        }


    }
}
