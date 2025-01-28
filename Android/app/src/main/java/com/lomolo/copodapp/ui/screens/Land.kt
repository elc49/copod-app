package com.lomolo.copodapp.ui.screens

import androidx.compose.foundation.layout.Arrangement
import androidx.compose.foundation.layout.Column
import androidx.compose.foundation.layout.Row
import androidx.compose.foundation.layout.Spacer
import androidx.compose.foundation.layout.fillMaxSize
import androidx.compose.foundation.layout.fillMaxWidth
import androidx.compose.foundation.layout.height
import androidx.compose.foundation.layout.padding
import androidx.compose.foundation.layout.size
import androidx.compose.foundation.lazy.LazyColumn
import androidx.compose.foundation.lazy.items
import androidx.compose.material.icons.Icons
import androidx.compose.material.icons.automirrored.twotone.ArrowBack
import androidx.compose.material3.CircularProgressIndicator
import androidx.compose.material3.ExperimentalMaterial3Api
import androidx.compose.material3.Icon
import androidx.compose.material3.IconButton
import androidx.compose.material3.MaterialTheme
import androidx.compose.material3.OutlinedCard
import androidx.compose.material3.Scaffold
import androidx.compose.material3.Surface
import androidx.compose.material3.Text
import androidx.compose.runtime.Composable
import androidx.compose.runtime.LaunchedEffect
import androidx.compose.runtime.collectAsState
import androidx.compose.runtime.getValue
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
import com.lomolo.copodapp.GetUserLandQuery
import com.lomolo.copodapp.R
import com.lomolo.copodapp.state.viewmodels.GetUserLands
import com.lomolo.copodapp.state.viewmodels.LandViewModel
import com.lomolo.copodapp.ui.common.TopBar
import com.lomolo.copodapp.ui.navigation.Navigation
import com.web3auth.core.types.UserInfo
import org.koin.androidx.compose.koinViewModel

object LandScreenDestination : Navigation {
    override val title = null
    override val route = "land"
}

@OptIn(ExperimentalMaterial3Api::class)
@Composable
fun LandScreen(
    modifier: Modifier = Modifier,
    viewModel: LandViewModel = koinViewModel<LandViewModel>(),
    userInfo: UserInfo?,
    onGoBack: () -> Unit,
) {
    LaunchedEffect(Unit) {
        viewModel.getUserLands(userInfo?.email!!)
    }
    val lands by viewModel.lands.collectAsState()
    val loading = when {
        viewModel.gettingUserLands is GetUserLands.Loading -> true
        else -> false
    }

    Scaffold(topBar = {
        TopBar(
            title = {
                Text(stringResource(R.string.your_lands))
            },
            navigationIcon = {
                IconButton(
                    onClick = onGoBack
                ) {
                    Icon(
                        Icons.AutoMirrored.TwoTone.ArrowBack,
                        contentDescription = stringResource(R.string.go_back)
                    )
                }
            }
        )
    }) { innerPadding ->
        Surface(
            modifier = modifier.padding(innerPadding)
        ) {
            when {
                viewModel.gettingUserLands is GetUserLands.Success -> {
                    if (lands.isEmpty()) {
                        NoLand()
                    } else {
                        LazyColumn(
                            verticalArrangement = Arrangement.spacedBy(8.dp),
                        ) {
                            items(lands) {
                                LandCard(land = it)
                            }
                        }
                    }
                }

                loading -> Column(
                    Modifier.fillMaxSize(),
                    verticalArrangement = Arrangement.Center,
                    horizontalAlignment = Alignment.CenterHorizontally,
                ) {
                    CircularProgressIndicator(
                        Modifier.size(20.dp)
                    )
                }

                viewModel.gettingUserLands is GetUserLands.Error -> ErrorScreen()
            }
        }
    }
}

@Composable
private fun LandCard(
    modifier: Modifier = Modifier,
    land: GetUserLandQuery.GetUserLand,
) {
    OutlinedCard(
        modifier = modifier
            .height(280.dp)
            .fillMaxWidth()
            .padding(8.dp),
    ) {
        AsyncImage(
            model = ImageRequest.Builder(LocalContext.current).data(land.url).crossfade(true)
                .build(),
            placeholder = painterResource(R.drawable.loading_img),
            error = painterResource(R.drawable.ic_broken_image),
            modifier = Modifier
                .fillMaxSize()
                .clip(MaterialTheme.shapes.extraSmall),
            contentScale = ContentScale.Crop,
            contentDescription = stringResource(R.string.land_title),
        )
    }
}

@Composable
private fun NoLand(
    modifier: Modifier = Modifier,
) {
    Column(
        modifier
            .fillMaxSize()
            .padding(8.dp),
        verticalArrangement = Arrangement.spacedBy(8.dp),
    ) {
        Text(
            stringResource(R.string.no_listings_yet),
            style = MaterialTheme.typography.titleLarge,
            textAlign = TextAlign.Start,
        )
        Text(stringResource(R.string.we_are_adding_listings))
        Row(
            modifier = Modifier.fillMaxWidth(), verticalAlignment = Alignment.CenterVertically
        ) {
            AsyncImage(
                model = ImageRequest.Builder(LocalContext.current).data(R.drawable._9872287)
                    .crossfade(true).build(),
                modifier = Modifier
                    .clip(MaterialTheme.shapes.medium)
                    .size(100.dp),
                placeholder = painterResource(R.drawable.loading_img),
                error = painterResource(R.drawable.ic_broken_image),
                contentDescription = stringResource(R.string.land)
            )
            Spacer(Modifier.size(16.dp))
            Column(
                verticalArrangement = Arrangement.Center,
            ) {
                Text(
                    stringResource(R.string.no_lands_available),
                    style = MaterialTheme.typography.titleMedium,
                )
                Text(
                    stringResource(R.string.you_can_create_a_listing),
                    style = MaterialTheme.typography.bodyMedium,
                    color = MaterialTheme.colorScheme.secondary,
                )
            }
        }
    }
}