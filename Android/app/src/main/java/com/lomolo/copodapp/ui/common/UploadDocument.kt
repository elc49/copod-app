package com.lomolo.copodapp.ui.common

import androidx.compose.foundation.clickable
import androidx.compose.foundation.layout.Arrangement
import androidx.compose.foundation.layout.Box
import androidx.compose.foundation.layout.Column
import androidx.compose.foundation.layout.PaddingValues
import androidx.compose.foundation.layout.RowScope
import androidx.compose.foundation.layout.fillMaxSize
import androidx.compose.foundation.layout.fillMaxWidth
import androidx.compose.foundation.layout.height
import androidx.compose.foundation.layout.padding
import androidx.compose.foundation.layout.size
import androidx.compose.material.icons.Icons
import androidx.compose.material.icons.automirrored.twotone.ArrowBack
import androidx.compose.material3.BottomAppBar
import androidx.compose.material3.Button
import androidx.compose.material3.CircularProgressIndicator
import androidx.compose.material3.ExperimentalMaterial3Api
import androidx.compose.material3.Icon
import androidx.compose.material3.IconButton
import androidx.compose.material3.MaterialTheme
import androidx.compose.material3.OutlinedCard
import androidx.compose.material3.Scaffold
import androidx.compose.material3.Text
import androidx.compose.material3.TopAppBar
import androidx.compose.runtime.Composable
import androidx.compose.ui.Alignment
import androidx.compose.ui.Modifier
import androidx.compose.ui.draw.clip
import androidx.compose.ui.layout.ContentScale
import androidx.compose.ui.platform.LocalContext
import androidx.compose.ui.res.painterResource
import androidx.compose.ui.res.stringResource
import androidx.compose.ui.unit.dp
import coil.compose.AsyncImage
import coil.request.ImageRequest
import com.lomolo.copodapp.R

@OptIn(ExperimentalMaterial3Api::class)
@Composable
fun UploadDocument(
    modifier: Modifier = Modifier,
    onNext: () -> Unit,
    title: @Composable (() -> Unit),
    onGoBack: () -> Unit,
    onSelectImage: () -> Unit,
    savingDoc: Boolean,
    image: Any,
    newUpload: Boolean = true,
    buttonText: @Composable (RowScope.() -> Unit),
) {
    val context = LocalContext.current

    Scaffold(topBar = {
        TopAppBar(title = title, navigationIcon = {
            IconButton(
                onClick = onGoBack,
            ) {
                Icon(
                    Icons.AutoMirrored.TwoTone.ArrowBack,
                    contentDescription = stringResource(R.string.go_back),
                )
            }
        })
    }, bottomBar = {
        BottomAppBar {
            Button(
                onClick = { if (!savingDoc) onNext() },
                shape = MaterialTheme.shapes.extraSmall,
                contentPadding = PaddingValues(16.dp),
                modifier = Modifier.fillMaxWidth(),
            ) {
                when (savingDoc) {
                    true -> CircularProgressIndicator(
                        Modifier.size(20.dp),
                        MaterialTheme.colorScheme.onPrimary,
                    )

                    false -> buttonText()
                }
            }
        }
    }) { innerPadding ->
        Box(
            modifier
                .fillMaxSize()
                .padding(innerPadding)
                .padding(8.dp),
        ) {
            Column(
                Modifier.padding(8.dp), verticalArrangement = Arrangement.spacedBy(8.dp)
            ) {
                Box(
                    contentAlignment = Alignment.Center,
                ) {
                    if (newUpload) {
                        OutlinedCard(
                            modifier = Modifier
                                .fillMaxWidth()
                                .height(360.dp),
                            onClick = onSelectImage,
                        ) {
                            Box(
                                Modifier
                                    .fillMaxSize()
                                    .padding(12.dp)
                            ) {
                                Icon(
                                    painterResource(R.drawable.doc_paper),
                                    modifier = Modifier.size(32.dp).align(Alignment.Center),
                                    contentDescription = stringResource(R.string.upload)
                                )
                                Text(
                                    stringResource(R.string.click_to_upload),
                                    style = MaterialTheme.typography.titleMedium,
                                    modifier = Modifier.align(Alignment.TopCenter),
                                )
                            }
                        }
                    } else {
                        AsyncImage(
                            model = ImageRequest.Builder(context).data(image).crossfade(true)
                                .build(),
                            contentScale = ContentScale.Fit,
                            placeholder = painterResource(R.drawable.loading_img),
                            error = painterResource(R.drawable.ic_broken_image),
                            modifier = Modifier
                                .fillMaxSize()
                                .clip(MaterialTheme.shapes.small)
                                .clickable { onSelectImage() },
                            contentDescription = stringResource(R.string.image),
                        )
                    }
                }
            }
        }
    }
}