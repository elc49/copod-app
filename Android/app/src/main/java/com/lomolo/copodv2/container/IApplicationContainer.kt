package com.lomolo.copodv2.container

import android.content.Context

interface IApplicationContainer {
}

class ApplicationContainer(private val context: Context): IApplicationContainer {}