package com.lomolo.copodapp.ui.screens

import androidx.compose.foundation.background
import androidx.compose.foundation.layout.Arrangement
import androidx.compose.foundation.layout.Box
import androidx.compose.foundation.layout.Column
import androidx.compose.foundation.layout.PaddingValues
import androidx.compose.foundation.layout.Row
import androidx.compose.foundation.layout.Spacer
import androidx.compose.foundation.layout.fillMaxSize
import androidx.compose.foundation.layout.fillMaxWidth
import androidx.compose.foundation.layout.padding
import androidx.compose.foundation.layout.size
import androidx.compose.material.icons.Icons
import androidx.compose.material.icons.twotone.Add
import androidx.compose.material.icons.twotone.Check
import androidx.compose.material3.Button
import androidx.compose.material3.Icon
import androidx.compose.material3.ListItem
import androidx.compose.material3.ListItemDefaults
import androidx.compose.material3.MaterialTheme
import androidx.compose.material3.Scaffold
import androidx.compose.material3.Surface
import androidx.compose.material3.Text
import androidx.compose.material3.TextButton
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
import com.lomolo.copodapp.GetOnboardingByEmailAndVerificationQuery
import com.lomolo.copodapp.R
import com.lomolo.copodapp.state.viewmodels.GetCurrentOnboarding
import com.lomolo.copodapp.state.viewmodels.MainViewModel
import com.lomolo.copodapp.state.viewmodels.OnboardingViewModel
import com.lomolo.copodapp.type.Verification
import com.lomolo.copodapp.ui.common.BottomNavBar
import com.lomolo.copodapp.ui.common.TopBar
import com.lomolo.copodapp.ui.navigation.Navigation
import okhttp3.internal.toImmutableList

object CreateLandScreenDestination : Navigation {
    override val title = null
    override val route = "create_land"
}

@Composable
fun CreateLandScreen(
    modifier: Modifier = Modifier,
    viewModel: OnboardingViewModel,
    mainViewModel: MainViewModel,
    onNavigateTo: (String) -> Unit,
    currentDestination: NavDestination,
    onClickAddLand: () -> Unit,
    onNext: (String) -> Unit,
) {
    LaunchedEffect(Unit) {
        viewModel.getCurrentOnboarding()
    }

    val currentOnboarding by viewModel.currentOnboarding.collectAsState()
    val reSubmits: List<String>? = currentOnboarding?.whoIsStillOnboarding()
    val showTopBar =
        viewModel.gettingCurrentOnboarding !is GetCurrentOnboarding.Loading && viewModel.gettingCurrentOnboarding !is GetCurrentOnboarding.Error

    Scaffold(bottomBar = {
        BottomNavBar(currentDestination = currentDestination, onNavigateTo = onNavigateTo)
    }, topBar = {
        TopBar(onClickAvatar = { onNext(AccountScreenDestination.route) },
            mainViewModel = mainViewModel,
            title = {
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
                        NoYourListings(innerPadding = innerPadding, onClickAddLand = onClickAddLand)
                    } else {
                        UnderReview(
                            innerPadding = innerPadding, onNext = onNext, reSubmits = reSubmits
                        )
                    }
                }

                GetCurrentOnboarding.Loading -> LoadingScreen()
                is GetCurrentOnboarding.Error -> ErrorScreen()
            }
        }
    }
}

@Composable
fun UnderReview(
    modifier: Modifier = Modifier,
    innerPadding: PaddingValues,
    reSubmits: List<String>?,
    onNext: (String) -> Unit,
) {
    Column(
        modifier
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
                    Icons.TwoTone.Check, contentDescription = stringResource(R.string.check)
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
        Spacer(Modifier.size(8.dp))
        if (reSubmits != null && reSubmits.isNotEmpty()) {
            Text(
                stringResource(R.string.resubmits),
                style = MaterialTheme.typography.titleMedium,
            )
            Spacer(Modifier.size(4.dp))
            reSubmits.map {
                ListItem(headlineContent = { Text(it) }, leadingContent = {
                    Box(
                        Modifier.padding(8.dp)
                    ) {
                        Icon(
                            painterResource(R.drawable.doc_paper),
                            modifier = Modifier
                                .size(28.dp)
                                .align(Alignment.Center),
                            contentDescription = stringResource(R.string.upload),
                        )
                    }
                }, colors = ListItemDefaults.colors(
                    containerColor = MaterialTheme.colorScheme.secondaryContainer,
                ), trailingContent = {
                    TextButton(
                        onClick = {
                            when (it) {
                                "Title" -> onNext("${UploadLandTitleScreenDestination.route}/${true}")
                                "SupportingDoc" -> onNext("${UploadGovtIssuedIdScreenDestination.route}/${true}")
                                "DisplayPicture" -> onNext("${UploadDisplayPictureDestination.route}/${true}")
                            }
                        },
                    ) {
                        Text(
                            stringResource(R.string.submit),
                            style = MaterialTheme.typography.titleMedium,
                        )
                    }
                })
            }
        }
    }
}

@Composable
fun NoYourListings(
    modifier: Modifier = Modifier,
    innerPadding: PaddingValues,
    onClickAddLand: () -> Unit,
) {
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
            model = ImageRequest.Builder(LocalContext.current).data(R.drawable._9872287)
                .crossfade(true).build(),
            modifier = Modifier.clip(MaterialTheme.shapes.medium),
            placeholder = painterResource(R.drawable.loading_img),
            error = painterResource(R.drawable.ic_broken_image),
            contentDescription = stringResource(R.string.land)
        )
        Button(
            onClick = onClickAddLand,
            contentPadding = PaddingValues(16.dp),
            modifier = Modifier.fillMaxWidth(),
        ) {
            Icon(
                Icons.TwoTone.Add,
                modifier = Modifier.size(24.dp),
                contentDescription = stringResource(R.string.add),
            )
            Spacer(Modifier.size(8.dp))
            Text(
                stringResource(R.string.start_onboarding),
                style = MaterialTheme.typography.titleMedium,
            )
        }
    }
}

fun GetOnboardingByEmailAndVerificationQuery.GetOnboardingByEmailAndVerification.whoIsStillOnboarding(): List<String> {
    var rejects = mutableListOf<String>()
    if (this.title.verified == Verification.REJECTED) {
        rejects.add(this.title.__typename.toString())
    }

    if (this.supportingDoc.verified == Verification.REJECTED) {
        rejects.add(this.supportingDoc.__typename.toString())
    }

    if (this.displayPicture.verified == Verification.REJECTED) {
        rejects.add(this.displayPicture.__typename.toString())
    }

    return rejects.toImmutableList()
}