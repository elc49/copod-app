package com.lomolo.copodapp.ui.screens

import androidx.compose.foundation.background
import androidx.compose.foundation.layout.Arrangement
import androidx.compose.foundation.layout.Box
import androidx.compose.foundation.layout.Column
import androidx.compose.foundation.layout.PaddingValues
import androidx.compose.foundation.layout.Row
import androidx.compose.foundation.layout.Spacer
import androidx.compose.foundation.layout.fillMaxSize
import androidx.compose.foundation.layout.padding
import androidx.compose.foundation.layout.size
import androidx.compose.material.icons.Icons
import androidx.compose.material.icons.twotone.Add
import androidx.compose.material.icons.twotone.Check
import androidx.compose.material3.Button
import androidx.compose.material3.Icon
import androidx.compose.material3.MaterialTheme
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
import androidx.compose.ui.platform.LocalContext
import androidx.compose.ui.res.painterResource
import androidx.compose.ui.res.stringResource
import androidx.compose.ui.unit.dp
import androidx.navigation.NavDestination
import coil.compose.AsyncImage
import coil.request.ImageRequest
import com.lomolo.copodapp.R
import com.lomolo.copodapp.state.viewmodels.GetCurrentOnboarding
import com.lomolo.copodapp.state.viewmodels.OnboardingViewModel
import com.lomolo.copodapp.ui.common.BottomNavBar
import com.lomolo.copodapp.ui.common.TopBar
import com.lomolo.copodapp.ui.navigation.Navigation

object AddLandScreenDestination : Navigation {
    override val title = null
    override val route = "add_land"
}

@Composable
fun AddLandScreen(
    modifier: Modifier = Modifier,
    viewModel: OnboardingViewModel,
    onNavigateTo: (String) -> Unit,
    currentDestination: NavDestination,
    onClickAddLand: () -> Unit,
) {
    LaunchedEffect(Unit) {
        viewModel.getCurrentOnboarding()
    }

    val currentOnboarding by viewModel.currentOnboarding.collectAsState()
    val showTopBar =
        viewModel.gettingCurrentOnboarding !is GetCurrentOnboarding.Loading && viewModel.gettingCurrentOnboarding !is GetCurrentOnboarding.Error

    Scaffold(bottomBar = {
        BottomNavBar(currentDestination = currentDestination, onNavigateTo = onNavigateTo)
    }, topBar = {
        TopBar(title = {
            if (showTopBar) {
                if (currentOnboarding == null) {
                    Text(stringResource(R.string.create_new_land))
                } else {
                    Text(stringResource(R.string.registration_status))
                }
            }
        })
    }) { innerPadding ->
        Surface(
            modifier = modifier.fillMaxSize()
        ) {
            when (viewModel.gettingCurrentOnboarding) {
                GetCurrentOnboarding.Success -> {
                    if (currentOnboarding == null) {
                        Column(
                            modifier
                                .fillMaxSize()
                                .padding(innerPadding)
                                .padding(8.dp),
                            verticalArrangement = Arrangement.spacedBy(8.dp),
                            horizontalAlignment = Alignment.CenterHorizontally,
                        ) {
                            Text(stringResource(R.string.new_land_copy))
                            AsyncImage(
                                model = ImageRequest.Builder(LocalContext.current)
                                    .data(R.drawable._9872287).crossfade(true).build(),
                                modifier = Modifier.clip(MaterialTheme.shapes.medium),
                                placeholder = painterResource(R.drawable.loading_img),
                                error = painterResource(R.drawable.ic_broken_image),
                                contentDescription = stringResource(R.string.land)
                            )
                            Button(
                                onClick = onClickAddLand,
                                contentPadding = PaddingValues(16.dp),
                                shape = MaterialTheme.shapes.extraSmall,
                                modifier = Modifier.align(Alignment.End),
                            ) {
                                Icon(
                                    Icons.TwoTone.Add,
                                    contentDescription = stringResource(R.string.add),
                                )
                                Spacer(Modifier.size(8.dp))
                                Text(
                                    stringResource(R.string.start_onboarding),
                                    style = MaterialTheme.typography.titleMedium,
                                )
                            }
                        }
                    } else {
                        Column(
                            Modifier
                                .fillMaxSize()
                                .padding(innerPadding)
                                .padding(8.dp),
                            verticalArrangement = Arrangement.spacedBy(4.dp)
                        ) {
                            Text(
                                stringResource(R.string.your_registration_status),
                                style = MaterialTheme.typography.titleLarge,
                            )
                            Row(
                                horizontalArrangement = Arrangement.spacedBy(16.dp),
                                verticalAlignment = Alignment.CenterVertically,
                            ) {
                                Box(
                                    Modifier
                                        .background(
                                            MaterialTheme.colorScheme.secondaryContainer,
                                            MaterialTheme.shapes.extraSmall,
                                        )
                                        .padding(16.dp),
                                    contentAlignment = Alignment.Center,
                                ) {
                                    Icon(
                                        Icons.TwoTone.Check,
                                        contentDescription = stringResource(R.string.check)
                                    )
                                }
                                Column {
                                    Text(
                                        stringResource(R.string.pending_review),
                                        style = MaterialTheme.typography.titleMedium,
                                    )
                                    Text(
                                        stringResource(R.string.review_copy)
                                    )
                                }
                            }
                        }
                    }
                }

                GetCurrentOnboarding.Loading -> LoadingScreen()
                is GetCurrentOnboarding.Error -> ErrorScreen()
            }
        }
    }
}
