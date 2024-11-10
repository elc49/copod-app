package com.lomolo.copodapp.util

import android.util.Log
import com.google.i18n.phonenumbers.PhoneNumberUtil
import com.google.i18n.phonenumbers.PhoneNumberUtil.PhoneNumberFormat
import com.google.i18n.phonenumbers.Phonenumber.PhoneNumber
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