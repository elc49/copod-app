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
    userInfo: UserInfo? = null,
    onOpenDialog: () -> Unit = {},
    title: @Composable (() -> Unit) = {},
    navigationIcon: @Composable (() -> Unit) = {}
) {
    TopAppBar(modifier = modifier, title = title, actions = {
        if (userInfo != null) {
            Avatar(
                avatar = userInfo.profileImage,
                name = userInfo.name,
                onClick = onOpenDialog,
            )
        }
    }, navigationIcon = navigationIcon)
}