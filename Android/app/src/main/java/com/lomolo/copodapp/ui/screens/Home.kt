package com.lomolo.copodapp.ui.screens

import androidx.compose.foundation.layout.Box
import androidx.compose.foundation.layout.PaddingValues
import androidx.compose.foundation.layout.fillMaxSize
import androidx.compose.foundation.layout.fillMaxWidth
import androidx.compose.foundation.layout.padding
import androidx.compose.material3.Button
import androidx.compose.material3.MaterialTheme
import androidx.compose.material3.Surface
import androidx.compose.material3.Text
import androidx.compose.runtime.Composable
import androidx.compose.ui.Alignment
import androidx.compose.ui.Modifier
import androidx.compose.ui.res.stringResource
import androidx.compose.ui.text.font.FontWeight
import androidx.compose.ui.unit.dp
import com.lomolo.copodapp.R
import com.lomolo.copodapp.ui.navigation.Navigation

object HomeScreenDestination : Navigation {
    override val title = null
    override val route = "home"
}

@Composable
fun HomeScreen(
    modifier: Modifier = Modifier,
    onGoToLogin: () -> Unit,
) {
    Surface {
        Box(
            modifier = modifier
                .fillMaxSize()
                .padding(8.dp)
        ) {
            Box(
                modifier = Modifier.align(Alignment.Center),
            ) {
                Text(
                    stringResource(R.string.home_headline),
                    style = MaterialTheme.typography.displayMedium,
                    fontWeight = FontWeight.ExtraBold,
                )
            }
            Button(
                onClick = onGoToLogin,
                modifier = Modifier
                    .fillMaxWidth()
                    .align(Alignment.BottomCenter),
                shape = MaterialTheme.shapes.extraSmall,
                contentPadding = PaddingValues(16.dp),
            ) {
                Text(
                    stringResource(R.string.get_started),
                    style = MaterialTheme.typography.titleMedium,
                )
            }
        }
    }
}