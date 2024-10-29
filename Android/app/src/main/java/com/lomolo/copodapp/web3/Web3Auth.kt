package com.lomolo.copodapp.web3

import android.content.Context
import android.net.Uri
import com.lomolo.copodapp.repository.IWeb3Auth
import com.web3auth.core.Web3Auth
import com.web3auth.core.types.Web3AuthOptions
import com.lomolo.copodapp.R
import com.web3auth.core.types.Network

fun getWeb3AuthImpl(context: Context): IWeb3Auth {
    val web3auth = Web3Auth(
        Web3AuthOptions(
            context = context,
            clientId = context.getString(R.string.web3auth_project_id),
            network = Network.SAPPHIRE_DEVNET,
            redirectUrl = Uri.parse("com.lomolo.copodapp://auth"),
        )
    )

    return com.lomolo.copodapp.repository.Web3Auth(web3auth)
}