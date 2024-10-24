package com.lomolo.copodv2

import android.app.Application
import com.lomolo.copodv2.container.ApplicationContainer
import com.lomolo.copodv2.container.IApplicationContainer

class CopodApp: Application() {
    lateinit var container: IApplicationContainer

    override fun onCreate() {
        super.onCreate()
        container = ApplicationContainer(this)
    }
}