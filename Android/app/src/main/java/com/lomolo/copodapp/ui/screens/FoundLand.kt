package com.lomolo.copodapp.ui.screens

import androidx.compose.foundation.clickable
import androidx.compose.foundation.layout.Arrangement
import androidx.compose.foundation.layout.Column
import androidx.compose.foundation.layout.fillMaxWidth
import androidx.compose.foundation.layout.padding
import androidx.compose.foundation.rememberScrollState
import androidx.compose.foundation.verticalScroll
import androidx.compose.material3.ListItem
import androidx.compose.material3.MaterialTheme
import androidx.compose.material3.Text
import androidx.compose.runtime.Composable
import androidx.compose.ui.Modifier
import androidx.compose.ui.res.stringResource
import androidx.compose.ui.unit.dp
import com.lomolo.copodapp.R
import com.lomolo.copodapp.ui.navigation.Navigation

object FoundLandScreenDestination : Navigation {
    override val title = null
    override val route = "found_land"
}

@Composable
fun FoundLandScreen(
    modifier: Modifier = Modifier,
) {
    val users = 100

    Column(
        modifier = modifier
            .padding(16.dp)
            .verticalScroll(rememberScrollState()),
        verticalArrangement = Arrangement.spacedBy(8.dp)
    ) {
        Text(
            stringResource(R.string.users_with_usage_rights),
            style = MaterialTheme.typography.titleLarge,
        )
        repeat(4) { idx ->
            val resultText = "User $idx"
            ListItem(headlineContent = {
                Text(
                    resultText,
                    style = MaterialTheme.typography.titleSmall,
                )
            }, supportingContent = {
                Text(
                    "email@email.com",
                    style = MaterialTheme.typography.bodyMedium,
                    color = MaterialTheme.colorScheme.secondary
                )
            }, modifier = Modifier
                .clickable {}
                .fillMaxWidth()
                .padding(vertical = 4.dp))
        }
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