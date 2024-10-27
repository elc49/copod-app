package com.lomolo.copodv2.util

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