package com.lomolo.copodapp.ui.common

import androidx.compose.foundation.background
import androidx.compose.foundation.clickable
import androidx.compose.foundation.layout.Box
import androidx.compose.foundation.layout.size
import androidx.compose.material3.MaterialTheme
import androidx.compose.material3.Text
import androidx.compose.runtime.Composable
import androidx.compose.ui.Alignment
import androidx.compose.ui.Modifier
import androidx.compose.ui.draw.clip
import androidx.compose.ui.layout.ContentScale
import androidx.compose.ui.platform.LocalContext
import androidx.compose.ui.res.stringResource
import androidx.compose.ui.unit.dp
import coil.compose.AsyncImage
import coil.request.ImageRequest
import com.lomolo.copodapp.R
import com.lomolo.copodapp.util.Util

@Composable
fun Avatar(
    avatar: String = "",
    email: String = "",
    onClick: () -> Unit = {},
) {
    if (avatar.isNotEmpty()) {
        AsyncImage(
            model = ImageRequest.Builder(LocalContext.current)
                .data(avatar)
                .crossfade(true),
            contentScale = ContentScale.Crop,
            modifier = Modifier
                .size(48.dp)
                .clip(MaterialTheme.shapes.extraSmall)
                .clickable { onClick() },
            contentDescription = stringResource(R.string.user),
        )
    } else {
        Box(
            Modifier
                .background(
                    MaterialTheme.colorScheme.surfaceVariant,
                    MaterialTheme.shapes.extraSmall
                )
                .size(48.dp)
                .clickable { onClick() },
            contentAlignment = Alignment.Center,
        ) {
            Text(
                Util.capitalize(email.first().toString())
            )
        }
    }
}