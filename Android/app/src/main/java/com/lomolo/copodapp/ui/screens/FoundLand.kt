package com.lomolo.copodapp.ui.screens

import androidx.compose.foundation.layout.Arrangement
import androidx.compose.foundation.layout.Column
import androidx.compose.foundation.layout.fillMaxSize
import androidx.compose.foundation.layout.fillMaxWidth
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
import com.lomolo.copodapp.ui.common.TopBar
import com.lomolo.copodapp.ui.navigation.Navigation

object FoundLandScreenDestination : Navigation {
    override val title = null
    override val route = "found_land"
}

@Composable
fun FoundLandScreen(
    modifier: Modifier = Modifier,
    onGoBack: () -> Unit,
) {
    Scaffold(topBar = {
        TopBar(title = { Text(stringResource(R.string.land_title_details)) }, navigationIcon = {
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
                    .padding(8.dp)
                    .verticalScroll(rememberScrollState()),
                verticalArrangement = Arrangement.spacedBy(12.dp)
            ) {
                // TODO style with box shadow
                Column(
                    Modifier
                        .fillMaxWidth()
                        .weight(.5f),
                ) {
                    Text(
                        stringResource(R.string.users_with_usage_rights),
                        style = MaterialTheme.typography.titleLarge,
                    )
                    Text("No users")
                }
                Column {
                    Text(
                        stringResource(R.string.land_registration_date),
                        style = MaterialTheme.typography.titleLarge,
                    )
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
                Column {
                    Text(
                        stringResource(R.string.previous_users),
                        style = MaterialTheme.typography.titleLarge,
                    )
                    Text(
                        stringResource(R.string.total, 0),
                        style = MaterialTheme.typography.titleSmall,
                    )
                    /*Text(
                        stringResource(R.string.had_usage_rights, 0),
                        style = MaterialTheme.typography.bodyMedium,
                        color = MaterialTheme.colorScheme.secondary,
                    )*/
                    Text(
                        stringResource(R.string.no_history),
                        style = MaterialTheme.typography.bodyMedium,
                        color = MaterialTheme.colorScheme.secondary,
                    )
                }
            }
        }
    }
}