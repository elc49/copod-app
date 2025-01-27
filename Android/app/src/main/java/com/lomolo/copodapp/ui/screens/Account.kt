package com.lomolo.copodapp.ui.screens

import androidx.compose.foundation.background
import androidx.compose.foundation.layout.Arrangement
import androidx.compose.foundation.layout.Box
import androidx.compose.foundation.layout.Column
import androidx.compose.foundation.layout.fillMaxSize
import androidx.compose.foundation.layout.fillMaxWidth
import androidx.compose.foundation.layout.padding
import androidx.compose.foundation.layout.size
import androidx.compose.foundation.shape.CircleShape
import androidx.compose.material.icons.Icons
import androidx.compose.material.icons.automirrored.twotone.ArrowBack
import androidx.compose.material.icons.automirrored.twotone.ArrowForward
import androidx.compose.material3.Button
import androidx.compose.material3.ButtonDefaults
import androidx.compose.material3.ExperimentalMaterial3Api
import androidx.compose.material3.Icon
import androidx.compose.material3.IconButton
import androidx.compose.material3.ListItem
import androidx.compose.material3.MaterialTheme
import androidx.compose.material3.Scaffold
import androidx.compose.material3.Surface
import androidx.compose.material3.Text
import androidx.compose.runtime.Composable
import androidx.compose.ui.Alignment
import androidx.compose.ui.Modifier
import androidx.compose.ui.draw.clip
import androidx.compose.ui.layout.ContentScale
import androidx.compose.ui.platform.LocalContext
import androidx.compose.ui.res.painterResource
import androidx.compose.ui.res.stringResource
import androidx.compose.ui.text.style.TextAlign
import androidx.compose.ui.unit.dp
import coil.compose.AsyncImage
import coil.request.ImageRequest
import com.lomolo.copodapp.R
import com.lomolo.copodapp.state.viewmodels.MainViewModel
import com.lomolo.copodapp.ui.common.TopBar
import com.lomolo.copodapp.ui.navigation.Navigation

object AccountScreenDestination : Navigation {
    override val title = null
    override val route = "account"
}

@OptIn(ExperimentalMaterial3Api::class)
@Composable
fun AccountScreen(
    modifier: Modifier = Modifier,
    mainViewModel: MainViewModel,
    onNext: (String) -> Unit,
    onGoBack: () -> Unit,
) {
    val userInfo = mainViewModel.userInfo

    Scaffold(topBar = {
        TopBar(title = {
            Text(stringResource(R.string.account))
        }, navigationIcon = {
            IconButton(
                onClick = onGoBack
            ) {
                Icon(
                    Icons.AutoMirrored.TwoTone.ArrowBack,
                    contentDescription = stringResource(R.string.go_back)
                )
            }
        })
    }) { innerPadding ->
        Surface(
            modifier = modifier.fillMaxSize()
        ) {
            Box(
                Modifier
                    .fillMaxSize()
                    .padding(innerPadding)
                    .padding(8.dp),
            ) {
                Box(
                    Modifier.align(Alignment.TopCenter),
                    contentAlignment = Alignment.Center,
                ) {
                    Column(
                        horizontalAlignment = Alignment.CenterHorizontally,
                        verticalArrangement = Arrangement.spacedBy(8.dp),
                    ) {
                        AsyncImage(
                            model = ImageRequest.Builder(LocalContext.current)
                                .data(userInfo?.profileImage).crossfade(true).build(),
                            modifier = Modifier
                                .size(120.dp)
                                .clip(CircleShape),
                            placeholder = painterResource(R.drawable.loading_img),
                            error = painterResource(R.drawable.ic_broken_image),
                            contentScale = ContentScale.Crop,
                            contentDescription = stringResource(R.string.account),
                        )
                        Text(
                            "${userInfo?.name}",
                            style = MaterialTheme.typography.titleLarge,
                            textAlign = TextAlign.Center,
                        )
                        Text(
                            "${userInfo?.email}",
                            textAlign = TextAlign.Center,
                        )
                    }
                }
                Box(
                    Modifier.align(Alignment.Center)
                ) {
                    Column {
                        ListItem(
                            headlineContent = { Text(stringResource(R.string.lands)) },
                            leadingContent = {
                                Box(
                                    Modifier
                                        .background(
                                            MaterialTheme.colorScheme.secondaryContainer,
                                            MaterialTheme.shapes.extraSmall,
                                        )
                                        .padding(8.dp),
                                    contentAlignment = Alignment.Center,
                                ) {
                                    Icon(
                                        painterResource(R.drawable.land),
                                        modifier = Modifier.size(24.dp),
                                        contentDescription = stringResource(R.string.land)
                                    )
                                }
                            },
                            trailingContent = {
                                IconButton(onClick = {
                                    onNext(LandScreenDestination.route)
                                }) {
                                    Icon(
                                        Icons.AutoMirrored.TwoTone.ArrowForward,
                                        contentDescription = stringResource(R.string.next)
                                    )
                                }
                            })
                        Button(
                            onClick = { mainViewModel.logOut() },
                            modifier = Modifier.fillMaxWidth().padding(16.dp),
                            colors = ButtonDefaults.buttonColors(
                                containerColor = MaterialTheme.colorScheme.secondaryContainer,
                                contentColor = MaterialTheme.colorScheme.primary,
                            ),
                        ) {
                            Text(
                                stringResource(R.string.sign_out),
                                style = MaterialTheme.typography.titleMedium,
                            )
                        }
                    }
                }
            }
        }
    }
}