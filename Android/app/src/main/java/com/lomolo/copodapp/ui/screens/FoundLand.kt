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
import androidx.compose.material3.MaterialTheme
import androidx.compose.material3.Scaffold
import androidx.compose.material3.Surface
import androidx.compose.material3.Text
import androidx.compose.runtime.Composable
import androidx.compose.ui.Modifier
import androidx.compose.ui.res.stringResource
import androidx.compose.ui.unit.dp
import com.lomolo.copodapp.R
import com.lomolo.copodapp.state.viewmodels.LandTitleDetailsViewModel
import com.lomolo.copodapp.ui.common.TopBar
import com.lomolo.copodapp.ui.navigation.Navigation
import org.koin.androidx.compose.koinViewModel

object FoundLandScreenDestination : Navigation {
    override val title = null
    override val route = "found_land"
    const val TITLE_NO_ARG = "titleNo"
    val routeWithArgs = "$route/{$TITLE_NO_ARG}"
}

@Composable
fun FoundLandScreen(
    modifier: Modifier = Modifier,
    onGoBack: () -> Unit,
    viewModel: LandTitleDetailsViewModel = koinViewModel<LandTitleDetailsViewModel>(),
) {
    val users = 100
    val titleNo = viewModel.titleNo

    Scaffold(topBar = {
        TopBar(title = { Text(titleNo) }, navigationIcon = {
            IconButton(onClick = onGoBack) {
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
            Column(
                modifier = Modifier
                    .fillMaxSize()
                    .padding(innerPadding)
                    .padding(16.dp)
                    .verticalScroll(rememberScrollState()),
                verticalArrangement = Arrangement.spacedBy(8.dp)
            ) {
                Text(
                    stringResource(R.string.users_with_usage_rights),
                    style = MaterialTheme.typography.titleLarge,
                )
                // TODO style with box shadow
                Text(
                    "No users"
                )
                Text(
                    stringResource(R.string.land_registration_date),
                    style = MaterialTheme.typography.titleLarge,
                )

                Column {
                    Text(
                        "January 1, 2024",
                        style = MaterialTheme.typography.titleSmall,
                    )
                    Text(
                        stringResource(R.string.land_registered_date),
                        style = MaterialTheme.typography.bodyMedium,
                        color = MaterialTheme.colorScheme.secondary,
                    )
                }

                Text(
                    stringResource(R.string.previous_users),
                    style = MaterialTheme.typography.titleLarge,
                )

                Column {
                    Text(
                        stringResource(R.string.total, users),
                        style = MaterialTheme.typography.titleSmall,
                    )
                    Text(
                        stringResource(R.string.had_usage_rights, users),
                        style = MaterialTheme.typography.bodyMedium,
                        color = MaterialTheme.colorScheme.secondary,
                    )
                }
            }
        }
    }
}