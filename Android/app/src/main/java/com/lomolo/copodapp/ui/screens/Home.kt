package com.lomolo.copodapp.ui.screens

import androidx.compose.foundation.layout.Box
import androidx.compose.foundation.layout.PaddingValues
import androidx.compose.foundation.layout.Row
import androidx.compose.foundation.layout.Spacer
import androidx.compose.foundation.layout.fillMaxSize
import androidx.compose.foundation.layout.fillMaxWidth
import androidx.compose.foundation.layout.padding
import androidx.compose.foundation.layout.size
import androidx.compose.material3.Button
import androidx.compose.material3.CircularProgressIndicator
import androidx.compose.material3.Icon
import androidx.compose.material3.MaterialTheme
import androidx.compose.material3.Surface
import androidx.compose.material3.Text
import androidx.compose.runtime.Composable
import androidx.compose.ui.Alignment
import androidx.compose.ui.Modifier
import androidx.compose.ui.res.painterResource
import androidx.compose.ui.res.stringResource
import androidx.compose.ui.text.font.FontWeight
import androidx.compose.ui.unit.dp
import com.lomolo.copodapp.R
import com.lomolo.copodapp.state.viewmodels.LoginSdk
import com.lomolo.copodapp.state.viewmodels.MainViewModel
import com.lomolo.copodapp.ui.navigation.Navigation

object HomeScreenDestination : Navigation {
    override val title = null
    override val route = "home"
}

@Composable
fun HomeScreen(
    modifier: Modifier = Modifier,
    mainViewModel: MainViewModel,
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
                onClick = { mainViewModel.login() },
                modifier = Modifier
                    .fillMaxWidth()
                    .align(Alignment.BottomCenter),
                shape = MaterialTheme.shapes.extraSmall,
                contentPadding = PaddingValues(16.dp),
            ) {
                when(mainViewModel.loginSdk) {
                    LoginSdk.Success -> Row(
                        verticalAlignment = Alignment.CenterVertically,
                    ) {
                        Icon(
                            painterResource(R.drawable.icons8_google),
                            modifier = Modifier.size(24.dp),
                            contentDescription = stringResource(R.string.google),
                        )
                        Spacer(Modifier.size(8.dp))
                        Text(
                            text = stringResource(R.string.google_sign_in),
                            style = MaterialTheme.typography.titleMedium,
                        )
                    }

                    LoginSdk.Loading -> CircularProgressIndicator(
                        Modifier.size(20.dp),
                        color = MaterialTheme.colorScheme.onPrimary,
                    )
                }
            }
        }
    }
}