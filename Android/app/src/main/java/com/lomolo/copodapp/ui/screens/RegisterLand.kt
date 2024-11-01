package com.lomolo.copodapp.ui.screens

import androidx.activity.compose.rememberLauncherForActivityResult
import androidx.activity.result.contract.ActivityResultContracts
import androidx.compose.foundation.Image
import androidx.compose.foundation.clickable
import androidx.compose.foundation.layout.Arrangement
import androidx.compose.foundation.layout.Box
import androidx.compose.foundation.layout.Column
import androidx.compose.foundation.layout.PaddingValues
import androidx.compose.foundation.layout.fillMaxSize
import androidx.compose.foundation.layout.fillMaxWidth
import androidx.compose.foundation.layout.padding
import androidx.compose.foundation.layout.size
import androidx.compose.foundation.rememberScrollState
import androidx.compose.foundation.verticalScroll
import androidx.compose.material.icons.Icons
import androidx.compose.material.icons.automirrored.twotone.ArrowBack
import androidx.compose.material3.Button
import androidx.compose.material3.ExperimentalMaterial3Api
import androidx.compose.material3.Icon
import androidx.compose.material3.IconButton
import androidx.compose.material3.MaterialTheme
import androidx.compose.material3.Scaffold
import androidx.compose.material3.Text
import androidx.compose.material3.TopAppBar
import androidx.compose.runtime.Composable
import androidx.compose.runtime.rememberCoroutineScope
import androidx.compose.ui.Modifier
import androidx.compose.ui.layout.ContentScale
import androidx.compose.ui.platform.LocalContext
import androidx.compose.ui.res.painterResource
import androidx.compose.ui.res.stringResource
import androidx.compose.ui.text.font.FontWeight
import androidx.compose.ui.unit.dp
import com.lomolo.copodapp.R
import com.lomolo.copodapp.ui.navigation.Navigation
import com.lomolo.copodapp.ui.viewmodels.RegisterLandViewModel
import org.koin.androidx.compose.koinViewModel

object RegisterLandScreenDestination : Navigation {
    override val title = R.string.register_land
    override val route = "register-land"
}

@OptIn(ExperimentalMaterial3Api::class)
@Composable
fun RegisterLandScreen(
    modifier: Modifier = Modifier,
    onGoBack: () -> Unit,
    viewModel: RegisterLandViewModel = koinViewModel<RegisterLandViewModel>(),
) {
    val context = LocalContext.current
    val scope = rememberCoroutineScope()
    val pickMedia = rememberLauncherForActivityResult(
        contract = ActivityResultContracts.PickVisualMedia()
    ) {
        if (it != null) {
            val stream = context.contentResolver.openInputStream(it)
            if (stream != null) {}
        }
    }

    Scaffold(topBar = {
        TopAppBar(title = {
            Text(stringResource(R.string.register_land))
        }, navigationIcon = {
            IconButton(
                onClick = onGoBack,
            ) {
                Icon(
                    Icons.AutoMirrored.TwoTone.ArrowBack,
                    contentDescription = stringResource(R.string.go_back),
                )
            }
        })
    }) { innerPadding ->
        Column(
            modifier
                .fillMaxSize()
                .padding(innerPadding)
                .padding(8.dp)
                .verticalScroll(rememberScrollState()),
            verticalArrangement = Arrangement.SpaceBetween,
        ) {
            Box {
                Column {
                    Text(
                        stringResource(R.string.land_title_document),
                        fontWeight = FontWeight.Bold,
                        style = MaterialTheme.typography.titleLarge,
                    )
                    Text(stringResource(R.string.govt_issued_title))
                    Image(
                        painterResource(R.drawable.upload),
                        contentScale = ContentScale.Fit,
                        modifier = Modifier
                            .fillMaxSize()
                            .size(240.dp)
                            .clickable{},
                        contentDescription = stringResource(R.string.image),
                    )
                }
            }
            Box {
                Column {
                    Text(
                        stringResource(R.string.govt_issued_id),
                        fontWeight = FontWeight.Bold,
                        style = MaterialTheme.typography.titleLarge,
                    )
                    Text(stringResource(R.string.upload_govt_issued_id))
                    Image(
                        painterResource(R.drawable.upload),
                        modifier = Modifier
                            .fillMaxSize()
                            .size(240.dp)
                            .clickable{},
                        contentScale = ContentScale.Fit,
                        contentDescription = stringResource(R.string.image),
                    )
                }
            }
            Button(
                onClick = {},
                shape = MaterialTheme.shapes.extraSmall,
                contentPadding = PaddingValues(16.dp),
                modifier = Modifier.fillMaxWidth(),
            ) {
                Text(
                    stringResource(R.string.submit),
                    style = MaterialTheme.typography.titleMedium,
                )
            }
        }
    }
}