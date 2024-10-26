package com.lomolo.copodv2.ui.screens

import androidx.compose.foundation.layout.Arrangement
import androidx.compose.foundation.layout.Column
import androidx.compose.foundation.layout.PaddingValues
import androidx.compose.foundation.layout.fillMaxSize
import androidx.compose.foundation.layout.fillMaxWidth
import androidx.compose.foundation.layout.padding
import androidx.compose.foundation.layout.size
import androidx.compose.foundation.text.KeyboardActions
import androidx.compose.foundation.text.KeyboardOptions
import androidx.compose.material.icons.Icons
import androidx.compose.material.icons.automirrored.twotone.ArrowBack
import androidx.compose.material3.Button
import androidx.compose.material3.CircularProgressIndicator
import androidx.compose.material3.ExperimentalMaterial3Api
import androidx.compose.material3.Icon
import androidx.compose.material3.IconButton
import androidx.compose.material3.MaterialTheme
import androidx.compose.material3.Scaffold
import androidx.compose.material3.Text
import androidx.compose.material3.TextField
import androidx.compose.material3.TopAppBar
import androidx.compose.runtime.Composable
import androidx.compose.runtime.collectAsState
import androidx.compose.runtime.getValue
import androidx.compose.ui.Modifier
import androidx.compose.ui.platform.LocalSoftwareKeyboardController
import androidx.compose.ui.res.stringResource
import androidx.compose.ui.text.input.ImeAction
import androidx.compose.ui.text.input.KeyboardType
import androidx.compose.ui.unit.dp
import com.lomolo.copodv2.R
import com.lomolo.copodv2.ui.navigation.Navigation
import com.lomolo.copodv2.viewmodels.LoginSdk
import com.lomolo.copodv2.viewmodels.MainViewModel

object LoginScreenDestination: Navigation {
    override val title = null
    override val route = "login"
}

@OptIn(ExperimentalMaterial3Api::class)
@Composable
fun LoginScreen(
    modifier: Modifier = Modifier,
    mainViewModel: MainViewModel,
    onGoBack: () -> Unit,
) {
    val loginData by mainViewModel.loginInput.collectAsState()
    val keyboardController = LocalSoftwareKeyboardController.current

    Scaffold(
        topBar = {
            TopAppBar(
                title = {Text(stringResource(R.string.start_by_email))},
                navigationIcon = {
                    IconButton(
                        onClick = onGoBack,
                    ) {
                        Icon(
                            Icons.AutoMirrored.TwoTone.ArrowBack,
                            contentDescription = stringResource(R.string.go_back),
                        )
                    }
                }
            )
        },
    ) { innerPadding ->
        Column(
            modifier = modifier
                .fillMaxSize()
                .padding(innerPadding),
            verticalArrangement = Arrangement.SpaceBetween,
        ) {
            TextField(
                value = loginData.email,
                onValueChange = { mainViewModel.setEmail(it) },
                modifier = Modifier.fillMaxWidth(),
                label = {
                    Text(stringResource(R.string.email))
                },
                placeholder = {
                    Text(stringResource(R.string.email))
                },
                keyboardOptions = KeyboardOptions(
                    keyboardType = KeyboardType.Text,
                    imeAction = ImeAction.Done,
                ),
                keyboardActions = KeyboardActions(
                    onDone = {
                        keyboardController?.hide()
                        mainViewModel.login()
                    },
                ),
            )
            Button(
                onClick = { mainViewModel.login() },
                modifier = Modifier.fillMaxWidth(),
                shape = MaterialTheme.shapes.extraSmall,
                contentPadding = PaddingValues(16.dp),
            ) {
                when(mainViewModel.loginSdk) {
                    LoginSdk.Success -> Text(
                        stringResource(R.string.signin),
                        style = MaterialTheme.typography.titleMedium,
                    )
                    LoginSdk.Loading -> CircularProgressIndicator(
                        Modifier.size(20.dp),
                        MaterialTheme.colorScheme.onPrimary,
                    )
                }
            }
        }
    }
}