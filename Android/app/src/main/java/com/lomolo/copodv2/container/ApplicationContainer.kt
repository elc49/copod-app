package com.lomolo.copodv2.container

import android.content.Context
import android.net.Uri
import com.lomolo.copodv2.R
import com.lomolo.copodv2.repository.IWeb3Auth
import com.web3auth.core.Web3Auth
import com.web3auth.core.types.BuildEnv
import com.web3auth.core.types.Network
import com.web3auth.core.types.Web3AuthOptions

interface IApplicationContainer {
    val web3Auth: IWeb3Auth
}

class ApplicationContainer(private val context: Context) : IApplicationContainer {
    private val web3auth = Web3Auth(
        Web3AuthOptions(
            context = context,
            clientId = context.getString(R.string.web3auth_project_id),
            network = Network.SAPPHIRE_DEVNET,
            buildEnv = BuildEnv.TESTING,
            redirectUrl = Uri.parse("com.lomolo.copodv2://auth"),
        )
    )

    override val web3Auth = com.lomolo.copodv2.repository.Web3Auth(web3auth)
}