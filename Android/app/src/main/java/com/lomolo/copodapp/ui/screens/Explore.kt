package com.lomolo.copodapp.ui.screens

import androidx.compose.foundation.layout.Arrangement
import androidx.compose.foundation.layout.Column
import androidx.compose.foundation.layout.fillMaxSize
import androidx.compose.foundation.layout.padding
import androidx.compose.foundation.layout.size
import androidx.compose.foundation.lazy.LazyColumn
import androidx.compose.material3.CircularProgressIndicator
import androidx.compose.material3.ExperimentalMaterial3Api
import androidx.compose.material3.MaterialTheme
import androidx.compose.material3.Scaffold
import androidx.compose.material3.Surface
import androidx.compose.material3.Text
import androidx.compose.runtime.Composable
import androidx.compose.runtime.collectAsState
import androidx.compose.runtime.getValue
import androidx.compose.ui.Alignment
import androidx.compose.ui.Modifier
import androidx.compose.ui.draw.clip
import androidx.compose.ui.platform.LocalContext
import androidx.compose.ui.res.painterResource
import androidx.compose.ui.res.stringResource
import androidx.compose.ui.text.font.FontWeight
import androidx.compose.ui.text.style.TextAlign
import androidx.compose.ui.unit.dp
import androidx.navigation.NavDestination
import coil.compose.AsyncImage
import coil.request.ImageRequest
import com.lomolo.copodapp.R
import com.lomolo.copodapp.state.viewmodels.GetLocalLands
import com.lomolo.copodapp.state.viewmodels.MarketViewModel
import com.lomolo.copodapp.ui.common.BottomNavBar
import com.lomolo.copodapp.ui.common.TopBar
import com.lomolo.copodapp.ui.navigation.Navigation
import org.koin.androidx.compose.koinViewModel

object ExploreMarketsScreenDestination : Navigation {
    override val title = null
    override val route = "explore"
}

@OptIn(ExperimentalMaterial3Api::class)
@Composable
fun ExploreMarketsScreen(
    modifier: Modifier = Modifier,
    marketViewModel: MarketViewModel = koinViewModel<MarketViewModel>(),
    onNavigateTo: (String) -> Unit,
    currentDestination: NavDestination,
) {
    val lands by marketViewModel.lands.collectAsState()

    Scaffold(topBar = {
        TopBar(
            title = {
                Text(stringResource(R.string.markets))
            }
        )
    }, bottomBar = {
        BottomNavBar(
            currentDestination = currentDestination, onNavigateTo = onNavigateTo
        )
    }) { innerPadding ->
        Surface(
            modifier = modifier.padding(innerPadding)
        ) {
            when (marketViewModel.gettingLands) {
                GetLocalLands.Success -> {
                    if (lands.isEmpty()) {
                        NoLands()
                    } else {
                        LazyColumn(
                            verticalArrangement = Arrangement.spacedBy(4.dp),
                        ) {
                            items(0) {
                                //LandCard(land = it)
                            }
                        }
                    }
                }

                GetLocalLands.Loading -> Column(
                    Modifier.fillMaxSize(),
                    verticalArrangement = Arrangement.Center,
                    horizontalAlignment = Alignment.CenterHorizontally,
                ) {
                    CircularProgressIndicator(
                        Modifier.size(20.dp)
                    )
                }

                is GetLocalLands.Error -> Text(
                    stringResource(R.string.something_wrong),
                    color = MaterialTheme.colorScheme.error,
                    fontWeight = FontWeight.Bold,
                )
            }
        }
    }
}


@Composable
private fun NoLands(
    modifier: Modifier = Modifier,
) {
    Column(
        modifier
            .fillMaxSize()
            .padding(8.dp),
        verticalArrangement = Arrangement.spacedBy(8.dp),
        horizontalAlignment = Alignment.CenterHorizontally,
    ) {
        AsyncImage(
            model = ImageRequest.Builder(LocalContext.current)
                .data(R.drawable._9872287)
                .crossfade(true)
                .build(),
            modifier = Modifier
                .clip(MaterialTheme.shapes.medium),
            placeholder = painterResource(R.drawable.loading_img),
            error = painterResource(R.drawable.ic_broken_image),
            contentDescription = stringResource(R.string.land)
        )
        Text(
            stringResource(R.string.no_land_space),
            style = MaterialTheme.typography.titleLarge,
            textAlign = TextAlign.Center,
        )
        Text(
            stringResource(R.string.check_market_later),
            textAlign = TextAlign.Center,
        )
    }
}