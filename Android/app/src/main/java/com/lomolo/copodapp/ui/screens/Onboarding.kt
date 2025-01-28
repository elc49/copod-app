package com.lomolo.copodapp.ui.screens

import androidx.compose.foundation.layout.Arrangement
import androidx.compose.foundation.layout.Column
import androidx.compose.foundation.layout.fillMaxSize
import androidx.compose.foundation.layout.padding
import androidx.compose.foundation.rememberScrollState
import androidx.compose.foundation.verticalScroll
import androidx.compose.material.icons.Icons
import androidx.compose.material.icons.automirrored.twotone.ArrowBack
import androidx.compose.material3.Icon
import androidx.compose.material3.IconButton
import androidx.compose.material3.Scaffold
import androidx.compose.material3.Surface
import androidx.compose.material3.Text
import androidx.compose.runtime.Composable
import androidx.compose.runtime.collectAsState
import androidx.compose.runtime.getValue
import androidx.compose.ui.Modifier
import androidx.compose.ui.res.stringResource
import androidx.compose.ui.unit.dp
import com.lomolo.copodapp.R
import com.lomolo.copodapp.state.viewmodels.OnboardingViewModel
import com.lomolo.copodapp.ui.OnboardDpCard
import com.lomolo.copodapp.ui.OnboardGovtIdCard
import com.lomolo.copodapp.ui.OnboardLandCard
import com.lomolo.copodapp.ui.common.TopBar
import com.lomolo.copodapp.ui.navigation.Navigation

object OnboardingScreenDestination : Navigation {
    override val title = R.string.onboarding
    override val route = "onboarding"
}

@Composable
fun OnboardingScreen(
    modifier: Modifier = Modifier,
    onGoBack: () -> Unit,
    onNext: (String) -> Unit,
    onboardingViewModel: OnboardingViewModel,
) {
    val land by onboardingViewModel.landTitle.collectAsState()
    val govtId by onboardingViewModel.supportingDoc.collectAsState()

    Scaffold(topBar = {
        TopBar(
            title = {
                Text(stringResource(R.string.onboarding))
            },
            navigationIcon = {
                IconButton(
                    onClick = onGoBack,
                ) {
                    Icon(
                        Icons.AutoMirrored.TwoTone.ArrowBack,
                        contentDescription = stringResource(R.string.go_back)
                    )
                }
            },
        )
    }) { innerPadding ->
        Surface(
            modifier = modifier.fillMaxSize()
        ) {
            Column(
                Modifier
                    .fillMaxSize()
                    .padding(innerPadding)
                    .padding(8.dp)
                    .verticalScroll(
                        rememberScrollState()
                    ),
                verticalArrangement = Arrangement.spacedBy(12.dp),
            ) {
                OnboardLandCard(onNext = onNext)
                OnboardGovtIdCard(onNext = onNext, land = land)
                OnboardDpCard(onNext =onNext, land = land, govtId = govtId)
            }
        }
    }
}





