package com.lomolo.copodapp.ui.common

import androidx.compose.foundation.clickable
import androidx.compose.foundation.layout.RowScope
import androidx.compose.foundation.layout.size
import androidx.compose.foundation.shape.CircleShape
import androidx.compose.material3.ExperimentalMaterial3Api
import androidx.compose.material3.TopAppBar
import androidx.compose.runtime.Composable
import androidx.compose.ui.Modifier
import androidx.compose.ui.draw.clip
import androidx.compose.ui.layout.ContentScale
import androidx.compose.ui.platform.LocalContext
import androidx.compose.ui.res.painterResource
import androidx.compose.ui.res.stringResource
import androidx.compose.ui.unit.dp
import coil.compose.AsyncImage
import coil.request.ImageRequest
import com.lomolo.copodapp.state.viewmodels.MainViewModel
import com.lomolo.copodapp.R

@OptIn(ExperimentalMaterial3Api::class)
@Composable
fun TopBar(
    modifier: Modifier = Modifier,
    title: @Composable (() -> Unit) = {},
    navigationIcon: @Composable (() -> Unit) = {},
    actions: @Composable (RowScope.() -> Unit) = {},
    mainViewModel: MainViewModel? = null,
    onClickAvatar: () -> Unit = {},
) {
    val userInfo = mainViewModel?.userInfo
    val actS: @Composable (RowScope.() -> Unit) = if (userInfo != null) {
        {
            actions
            AsyncImage(
                model = ImageRequest.Builder(LocalContext.current)
                    .data(userInfo.profileImage)
                    .crossfade(true)
                    .build(),
                modifier = Modifier
                    .size(32.dp)
                    .clip(CircleShape)
                    .clickable { onClickAvatar() },
                contentScale = ContentScale.Crop,
                placeholder = painterResource(R.drawable.loading_img),
                error = painterResource(R.drawable.ic_broken_image),
                contentDescription = stringResource(R.string.account),
            )
        }
    } else {
        actions
    }

    TopAppBar(
        modifier = modifier,
        actions = actS,
        title = title,
        navigationIcon = navigationIcon
    )
}