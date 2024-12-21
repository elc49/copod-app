package com.lomolo.copodapp.ui.screens

import androidx.compose.foundation.layout.Arrangement
import androidx.compose.foundation.layout.Column
import androidx.compose.foundation.layout.fillMaxSize
import androidx.compose.foundation.layout.fillMaxWidth
import androidx.compose.foundation.layout.height
import androidx.compose.foundation.layout.padding
import androidx.compose.foundation.layout.size
import androidx.compose.foundation.lazy.LazyColumn
import androidx.compose.foundation.lazy.items
import androidx.compose.material.icons.Icons
import androidx.compose.material.icons.twotone.Add
import androidx.compose.material3.CircularProgressIndicator
import androidx.compose.material3.ExperimentalMaterial3Api
import androidx.compose.material3.Icon
import androidx.compose.material3.MaterialTheme
import androidx.compose.material3.OutlinedCard
import androidx.compose.material3.OutlinedIconButton
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
import androidx.compose.ui.unit.dp
import androidx.navigation.NavDestination
import coil.compose.AsyncImage
import coil.request.ImageRequest
import com.lomolo.copodapp.GetOnboardingByEmailQuery
import com.lomolo.copodapp.GetUserLandQuery
import com.lomolo.copodapp.R
import com.lomolo.copodapp.state.viewmodels.GetCurrentOnboarding
import com.lomolo.copodapp.state.viewmodels.GetUserLands
import com.lomolo.copodapp.state.viewmodels.LandViewModel
import com.lomolo.copodapp.state.viewmodels.OnboardingViewModel
import com.lomolo.copodapp.type.Verification
import com.lomolo.copodapp.ui.common.BottomNavBar
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
    currentDestination: NavDestination,
    onNavigateTo: (String) -> Unit,
    viewModel: LandViewModel = koinViewModel<LandViewModel>(),
    onboardingViewModel: OnboardingViewModel,
    userInfo: UserInfo?,
    onClickAddLand: () -> Unit,
) {
    LaunchedEffect(Unit) {
        viewModel.getUserLands(userInfo?.email!!)
        onboardingViewModel.getCurrentOnboarding()
    }
    val currentOnboarding by onboardingViewModel.currentOnboarding.collectAsState()
    val lands by viewModel.lands.collectAsState()
    val loading = when {
        viewModel.gettingUserLands is GetUserLands.Loading -> true
        onboardingViewModel.gettingCurrentOnboarding is GetCurrentOnboarding.Loading -> true
        else -> false
    }

    Scaffold(topBar = {
        TopBar(
            title = {
                if (currentOnboarding?.isOnboardingOK() != true) {
                    Text(stringResource(R.string.verification_progress_text))
                } else {
                    Text(stringResource(R.string.your_lands))
                }
            },
        )
    }, bottomBar = {
        BottomNavBar(
            currentDestination = currentDestination,
            onNavigateTo = onNavigateTo,
        )
    }) { innerPadding ->
        Surface(
            modifier = modifier.padding(innerPadding)
        ) {
            when {
                currentOnboarding?.isOnboardingOK() != true -> Column(
                    Modifier.fillMaxSize(),
                    verticalArrangement = Arrangement.spacedBy(8.dp),
                    horizontalAlignment = Alignment.CenterHorizontally,
                ) {
                    AsyncImage(
                        model = ImageRequest.Builder(LocalContext.current)
                            .data(R.drawable.doc_review).crossfade(true).build(),
                        contentScale = ContentScale.Crop,
                        contentDescription = stringResource(R.string.waiting_submission)
                    )
                    Column(
                        Modifier.padding(16.dp), verticalArrangement = Arrangement.spacedBy(8.dp)
                    ) {
                        Text(
                            stringResource(R.string.under_review_text),
                            style = MaterialTheme.typography.titleLarge,
                        )
                        Text(
                            stringResource(R.string.verification_notice_text),
                        )
                    }
                }

                currentOnboarding?.isOnboardingOK() == true && viewModel.gettingUserLands is GetUserLands.Success -> {
                    if (lands.isEmpty()) {
                        NoLand(
                            onClickAddLand = onClickAddLand,
                        )
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

                // TODO: who is still onboarding
                // TODO: move to the screen

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
    onClickAddLand: () -> Unit,
) {
    Column(
        modifier.fillMaxSize(),
        verticalArrangement = Arrangement.Center,
        horizontalAlignment = Alignment.CenterHorizontally,
    ) {
        Icon(
            painterResource(R.drawable.sealed),
            modifier = Modifier.size(60.dp),
            contentDescription = stringResource(R.string.land),
        )
        Text(stringResource(R.string.no_user_lands))
        OutlinedIconButton(
            onClick = onClickAddLand
        ) {
            Icon(
                Icons.TwoTone.Add,
                contentDescription = stringResource(R.string.add),
            )
        }
    }
}

fun GetOnboardingByEmailQuery.GetOnboardingByEmail.isOnboardingOK(): Boolean {
    return when {
        this.displayPicture.verified == Verification.VERIFIED && this.supportingDoc.verified == Verification.VERIFIED && this.title.verified == Verification.VERIFIED -> true

        else -> false
    }
}

fun GetOnboardingByEmailQuery.GetOnboardingByEmail.whoIsStillOnboarding(): List<String> {
    var rejects = mutableListOf<String>()
    when {
        this.displayPicture.verified == Verification.REJECTED -> rejects.add(this.displayPicture.__typename.lowercase())
        this.title.verified == Verification.REJECTED -> rejects.add(this.title.__typename.lowercase())
        this.supportingDoc.verified == Verification.REJECTED -> rejects.add(this.supportingDoc.__typename.lowercase())
    }
    return rejects
}