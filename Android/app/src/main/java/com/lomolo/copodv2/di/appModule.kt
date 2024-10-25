package com.lomolo.copodv2.di

import android.content.Context
import android.net.Uri
import com.lomolo.copodv2.repository.IWeb3Auth
import com.lomolo.copodv2.viewmodels.MainViewModel
import com.web3auth.core.Web3Auth
import com.web3auth.core.types.Web3AuthOptions
import org.koin.core.module.dsl.viewModel
import org.koin.dsl.module
import com.lomolo.copodv2.R
import com.web3auth.core.types.BuildEnv
import com.web3auth.core.types.Network

val appModule = module {
    single {
        getWeb3AuthImpl(get())
    }

    viewModel { MainViewModel(get()) }
}

private fun getWeb3AuthImpl(context: Context): IWeb3Auth {
    val web3auth = Web3Auth(
        Web3AuthOptions(
            context = context,
            clientId = context.getString(R.string.web3auth_project_id),
            network = Network.SAPPHIRE_DEVNET,
            buildEnv = BuildEnv.TESTING,
            redirectUrl = Uri.parse("com.lomolo.copodv2://auth"),
        )
    )

    return com.lomolo.copodv2.repository.Web3Auth(web3auth)
}