package com.lomolo.copodapp.ui.common

import androidx.compose.material3.ExperimentalMaterial3Api
import androidx.compose.material3.TopAppBar
import androidx.compose.runtime.Composable
import androidx.compose.ui.Modifier
import com.web3auth.core.types.UserInfo

@OptIn(ExperimentalMaterial3Api::class)
@Composable
fun TopBar(
    modifier: Modifier = Modifier,
    userInfo: UserInfo,
    onOpenDialog: () -> Unit,
) {
    TopAppBar(title = {}, actions = {
        Avatar(
            avatar = userInfo.profileImage,
            email = userInfo.email,
            onClick = onOpenDialog,
        )
    })
}